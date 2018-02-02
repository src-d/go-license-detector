package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/src-d/go-license-detector.v1"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: licenseng <path>")
		os.Exit(1)
	}
	arg := os.Args[1]
	_, err := os.Stat(arg)
	var licenses map[string]float32
	if err == nil {
		licenses, err = getLicensesByPath(arg)
	} else {
		licenses, err = getLicensesByURL(arg)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	licenseStrs := make([]string, 0, len(licenses))
	for key, val := range licenses {
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

func getLicensesByPath(path string) (map[string]float32, error) {
	return ld.InvestigateProjectLicenses(path)
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
	return ld.InvestigateFilesLicenses(fileNames, func(file string) (string, error) {
		return files[file].Contents()
	})
}
