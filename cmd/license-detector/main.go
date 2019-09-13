// license-detector prints the most probable licenses for a repository
// given either its path in the local file system or a URL pointing to
// the repository.
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/pflag"
	"gopkg.in/src-d/go-license-detector.v3/licensedb"
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
	results := licensedb.Analyse(args...)

	switch format {
	case "text":
		for _, res := range results {
			fmt.Fprintln(writer, res.Arg)
			if res.ErrStr != "" {
				fmt.Fprintf(writer, "\t%v\n", res.ErrStr)
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
