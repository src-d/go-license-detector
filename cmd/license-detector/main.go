// license-detector prints the most probable licenses for a repository
// given either its path in the local file system or a URL pointing to
// the repository.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"sort"
	"sync"

	"github.com/spf13/pflag"
	"gopkg.in/src-d/go-license-detector.v2/licensedb"
	"gopkg.in/src-d/go-license-detector.v2/licensedb/filer"
)

func main() {
	format := pflag.StringP("format", "f", "text", "Output format: json, text")
	pflag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:  license-detector path ...")
		pflag.PrintDefaults()
	}
	pflag.Parse()
	if (*format != "json" && *format != "text") || pflag.NArg() == 0 {
		pflag.Usage()
		os.Exit(1)
	}
	detect(pflag.Args(), *format, os.Stdout)
}

// detect runs license analysis on each item in `args`` and outputs
// the results in the specified `format` to `writer`.
func detect(args []string, format string, writer io.Writer) {
	nargs := len(args)
	results := make([]result, nargs)
	var wg sync.WaitGroup
	wg.Add(nargs)
	for i, arg := range args {
		go func(i int, arg string) {
			defer wg.Done()
			matches, err := process(arg)
			res := result{Arg: arg, Matches: matches, Err: err, ErrStr: ""}
			if err != nil {
				res.ErrStr = err.Error()
			}
			results[i] = res
		}(i, arg)
	}
	wg.Wait()

	switch format {
	case "text":
		for _, res := range results {
			fmt.Fprintln(writer, res.Arg)
			if res.Err != nil {
				fmt.Fprintf(writer, "\t%v\n", res.Err)
				continue
			}
			for _, m := range res.Matches {
				fmt.Fprintf(writer, "\t%1.f%%\t%s\n", 100*m.Confidence, m.License)
			}
		}
	case "json":
		b, err := json.MarshalIndent(results, "", "\t")
		if err != nil {
			log.Fatalf("could not encode result to JSON: %v", err)
		}
		fmt.Fprintf(writer, "%s\n", b)
	}
}

// json cannot not marshal error-s as we would expect (we always get "{}")
// so we have to include ErrStr which is Err.Error()
type result struct {
	Arg     string  `json:"project,omitempty"`
	Matches []match `json:"matches,omitempty"`
	Err     error   `json:"-"`
	ErrStr  string  `json:"error,omitempty"`
}

type match struct {
	License    string  `json:"license"`
	Confidence float32 `json:"confidence"`
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
