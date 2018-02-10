package licensedb

import (
	"fmt"
	"regexp"
	"strings"

	"gopkg.in/src-d/go-license-detector.v1/licensedb/filer"
	"gopkg.in/src-d/go-license-detector.v1/licensedb/internal/processors"
)

var (
	licenses = loadLicenses()

	// Base names of guessable license files.
	fileNames = []string{
		"copying",
		"copyleft",
		"copyright",
		"license",
		"unlicense",
		"licence",
	}

	// License file extensions. Combined with the fileNames slice
	// to create a set of files we can reasonably assume contain
	// licensing information.
	fileExtensions = []string{
		"",
		".md",
		".rst",
		".html",
		".txt",
	}

	filePreprocessors = map[string]func([]byte) []byte{
		".md":   processors.Markdown,
		".rst":  processors.RestructuredText,
		".html": processors.HTML,
	}

	licenseFileRe = regexp.MustCompile(
		fmt.Sprintf("^(%s)(%s)$",
			strings.Join(fileNames, "|"),
			strings.Replace(strings.Join(fileExtensions, "|"), ".", "\\.", -1)))

	readmeFileRe = regexp.MustCompile(fmt.Sprintf("^readme(%s)$",
		strings.Replace(strings.Join(fileExtensions, "|"), ".", "\\.", -1)))
)

// Detect scans the given list of file names and detects the licenses.
// Each match has the confidence assigned, from 0 to 1, 1 means 100% confident.
func Detect(f filer.Filer) (map[string]float32, error) {
	patterns := []struct {
		regexp *regexp.Regexp
		query  func(string) map[string]float32
	}{
		{licenseFileRe, licenses.QueryLicenseText},
		{readmeFileRe, licenses.QueryReadmeText},
	}
	for _, pattern := range patterns {
		files, err := filer.Contents(filer.WithPathRegexp(f, pattern.regexp))
		if err != nil {
			return nil, err
		} else if len(files) == 0 {
			continue
		}

		ls := map[string]float32{}
		for _, file := range files {
			for name, sim := range pattern.query(string(file)) {
				if sim > ls[name] {
					ls[name] = sim
				}
			}
		}
		return ls, nil
	}
	return nil, nil
}
