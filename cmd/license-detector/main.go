// license-detector prints the most probable licenses for a repository
// given either its path in the local file system or a URL pointing to
// the repository.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"sort"
	"sync"

	"gopkg.in/src-d/go-license-detector.v1/detector"
	"gopkg.in/src-d/go-license-detector.v1/detector/filer"
)

func main() {
	format := flag.String("f", "text", "Output format: json, text")
	flag.Usage = func() {
		fmt.Println("Usage:  license-detector path ...")
		flag.PrintDefaults()
	}
	flag.Parse()
	if (*format != "json" && *format != "text") || flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	type result struct {
		Arg     string
		Matches []match `json:",omitempty"`
		Err     error   `json:",omitempty"`
	}
	results := make([]result, flag.NArg())
	var wg sync.WaitGroup
	wg.Add(flag.NArg())
	for i, arg := range flag.Args() {
		go func(i int, arg string) {
			defer wg.Done()
			matches, err := process(*format, arg)
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
				fmt.Printf("\t%1.f%%\t%s\n", 100*m.P, m.License)
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
	License string
	P       float32
}

func process(format, arg string) ([]match, error) {
	newFiler := filer.FromDirectory
	if _, err := url.Parse(arg); err == nil {
		newFiler = filer.FromGitURL
	}

	filer, err := newFiler(arg)
	if err != nil {
		return nil, err
	}

	ls, err := detector.Licenses(filer)
	if err != nil {
		return nil, err
	}

	var matches []match
	for k, v := range ls {
		matches = append(matches, match{k, v})
	}
	sort.Slice(matches, func(i, j int) bool { return matches[i].P > matches[j].P })
	return matches, nil
}
