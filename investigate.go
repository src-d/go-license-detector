package licenseng

import (
	"errors"
	"fmt"
	"io/ioutil"
	paths "path"
	"regexp"
	"strings"
)

var (
	NoLicenseFoundError   = errors.New("no license file was found")

	globalLicenseDatabase = &LicenseDatabase{}
	// Base names of guessable license files.
	fileNames = []string{
		"copying",
		"copyleft",
		"copyright",
		"license",
		"unlicense",
	}

	// License file extensions. Combined with the fileNames slice
	// to create a set of files we can reasonably assume contain
	// licensing information.
	fileExtensions = []string{
		"",
		".md",
		".rst",
		".txt",
	}

	filePreprocessors = map[string]func(string)string {
		".md": PreprocessMarkdown,
		".rst": PreprocessRestructuredText,
	}

	fileRe = regexp.MustCompile(
		fmt.Sprintf("^(%s)(%s)$",
			strings.Join(fileNames, "|"),
			strings.Replace(strings.Join(fileExtensions, "|"), ".", "\\.", -1)))
)

func InvestigateProjectLicenses(path string) (map[string]float32, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	candidates := []string{}
	for _, file := range files {
		if !file.IsDir() && fileRe.MatchString(strings.ToLower(file.Name())) {
			textBytes, err := ioutil.ReadFile(paths.Join(path, file.Name()))
			if err == nil {
				text := string(textBytes)
				if preprocessor, exists := filePreprocessors[paths.Ext(file.Name())]; exists {
					text = preprocessor(text)
				}
				candidates = append(candidates, text)
			}
		}
	}
	if len(candidates) == 0 {
		return nil, NoLicenseFoundError
	}
	maxLicenses := map[string]float32{}
	for _, text := range candidates {
		options, similarities := InvestigateFileLicense(text)
		for i, sim := range similarities {
			maxSim := maxLicenses[options[i]]
			if sim > maxSim {
				maxLicenses[options[i]] = sim
			}
		}
	}
	return maxLicenses, nil
}

func InvestigateFileLicense(text string) (options []string, similarities []float32) {
	return globalLicenseDatabase.Query(text)
}

func init() {
	globalLicenseDatabase.Load()
}
