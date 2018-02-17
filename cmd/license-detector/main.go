// license-detector prints the most probable licenses for a repository
// given either its path in the local file system or a URL pointing to
// the repository.
package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"encoding/json"
	"github.com/spf13/pflag"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/src-d/go-license-detector.v1/licensedb"
	"sync"
)

func main() {
	var outputFormat string
	pflag.StringVarP(&outputFormat, "format", "f", "text", "Output format, choose one of json, text.")
	pflag.Parse()
	if pflag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Usage: license-detector <path> [<path>...]")
		os.Exit(1)
	}
	formatter := formatters[outputFormat]
	if formatter == nil {
		fmt.Fprintf(os.Stderr, "Unknown format \"%s\", choose one of json, text.\n", outputFormat)
		os.Exit(1)
	}
	results := make([]analysisResult, pflag.NArg())
	wg := sync.WaitGroup{}
	wg.Add(pflag.NArg())
	for argIndex, arg := range pflag.Args() {
		go func(argIndex int, arg string) {
			defer wg.Done()
			_, err := os.Stat(arg)
			var licenses map[string]float32
			if err == nil {
				licenses, err = getLicensesByDirectory(arg)
			} else {
				licenses, err = getLicensesByURL(arg)
			}
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			results[argIndex] = analysisResult{Name: arg, Licenses: licenses}
		}(argIndex, arg)
	}
	wg.Wait()
	formatter(results)
}

type analysisResult struct {
	Name     string
	Licenses map[string]float32
}

var formatters = map[string]func([]analysisResult){
	"json": formatJSON,
	"text": formatText,
}

func formatJSON(results []analysisResult) {
	serialized, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(serialized)
	fmt.Println()
}

func formatText(results []analysisResult) {
	for resultIndex, result := range results {
		if len(results) > 1 {
			if resultIndex > 0 {
				fmt.Println()
			}
			fmt.Println(result.Name)
		}
		licenseStrs := make([]string, 0, len(result.Licenses))
		for key, val := range result.Licenses {
			licenseStrs = append(licenseStrs, fmt.Sprintf("%.1f%%\t%s", val*100, key))
		}
		sort.Sort(sort.Reverse(sort.StringSlice(licenseStrs)))
		// 100.0% is sadly but fairly placed at the end
		for _, s := range licenseStrs {
			if strings.HasPrefix(s, "100.0%") {
				fmt.Println(s)
			}
		}
		for _, s := range licenseStrs {
			if !strings.HasPrefix(s, "100.0%") {
				fmt.Println(s)
			}
		}
	}
}

func getLicensesByDirectory(path string) (map[string]float32, error) {
	return licensedb.InvestigateProjectLicenses(path)
}

func getLicensesByURL(url string) (map[string]float32, error) {
	cloneOptions := &git.CloneOptions{URL: url}
	repository, err := git.Clone(memory.NewStorage(), nil, cloneOptions)
	if err != nil {
		return nil, err
	}
	head, err := repository.Head()
	if err != nil {
		return nil, err
	}
	commit, err := repository.CommitObject(head.Hash())
	if err != nil {
		return nil, err
	}
	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}
	fileNames := []string{}
	files := map[string]*object.File{}
	tree.Files().ForEach(func(file *object.File) error {
		fileNames = append(fileNames, file.Name)
		files[file.Name] = file
		return nil
	})
	return licensedb.InvestigateFilesLicenses(fileNames, func(file string) (string, error) {
		return files[file].Contents()
	})
}
