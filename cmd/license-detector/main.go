// license-detector prints the most probable licenses for a repository
// given either its path in the local file system or a URL pointing to
// the repository.
package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sort"
	"sync"

	"github.com/spf13/pflag"
	"gopkg.in/src-d/go-license-detector.v2/licensedb"
	"gopkg.in/src-d/go-license-detector.v2/licensedb/filer"
)

func main() {
	format := pflag.String("f", "text", "Output format: json, text")
	pflag.Usage = func() {
		fmt.Println("Usage:  license-detector path ...")
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if (*format != "json" && *format != "text") || pflag.NArg() == 0 {
		pflag.Usage()
		os.Exit(1)
	}

	type result struct {
		Arg     string
		Matches []match `json:",omitempty"`
		Err     error   `json:",omitempty"`
	}
	results := make([]result, pflag.NArg())
	var wg sync.WaitGroup
	wg.Add(pflag.NArg())
	for i, arg := range pflag.Args() {
		go func(i int, arg string) {
			defer wg.Done()
			matches, err := process(arg)
			results[i] = result{arg, matches, err}
		}(i, arg)
	}
	wg.Wait()

	switch *format {
	case "text":
		for _, res := range results {
			fmt.Println(res.Arg)
			if res.Err != nil {
				fmt.Printf("\t%v\n", res.Err)
				continue
			}
			for _, m := range res.Matches {
				fmt.Printf("\t%1.f%%\t%s\n", 100*m.Confidence, m.License)
			}
		}
	case "json":
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not encode result to JSON: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}

type match struct {
	License    string
	Confidence float32
}

func process(arg string) ([]match, error) {
	newFiler := filer.FromDirectory
	fi, err := os.Stat(arg)
	if err != nil {
		if _, err := url.Parse(arg); err == nil {
			newFiler = filer.FromGitURL
		}
	} else if !fi.IsDir() {
		newFiler = filer.FromSiva
	}

	resolvedFiler, err := newFiler(arg)
	if err != nil {
		return nil, err
	}

	ls, err := licensedb.Detect(resolvedFiler)
	if err != nil {
		return nil, err
	}

	var matches []match
	for k, v := range ls {
		matches = append(matches, match{k, v})
	}
	sort.Slice(matches, func(i, j int) bool { return matches[i].Confidence > matches[j].Confidence })
	return matches, nil
}
