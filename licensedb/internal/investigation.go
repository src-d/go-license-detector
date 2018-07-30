package internal

import (
	"bytes"
	"fmt"
	paths "path"
	"regexp"
	"strings"
	"sync"

	"gopkg.in/src-d/go-license-detector.v2/licensedb/filer"
	"gopkg.in/src-d/go-license-detector.v2/licensedb/internal/processors"
	"gopkg.in/src-d/enry.v1"
)

var (
	globalLicenseDB struct {
		sync.Once
		*database
	}
	globalLicenseDatabase = func() *database {
		globalLicenseDB.Once.Do(func() {
			globalLicenseDB.database = loadLicenses()
		})
		return globalLicenseDB.database
	}

	// Base names of guessable license files
	licenseFileNames = []string{
		"li[cs]en[cs]e(s?)",
		"legal",
		"copy(left|right|ing)",
		"unlicense",
		"l?gpl([-_ v]?)(\\d\\.?\\d)?",
		"bsd",
		"mit",
		"apache",
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
		fmt.Sprintf("^(|.*[-_. ])(%s)(|[-_. ].*)$",
			strings.Join(licenseFileNames, "|")))

	readmeFileRe = regexp.MustCompile(fmt.Sprintf("^(readme|guidelines)(%s)$",
		strings.Replace(strings.Join(fileExtensions, "|"), ".", "\\.", -1)))

	licenseDirectoryRe = regexp.MustCompile(fmt.Sprintf(
		"^(%s)$", strings.Join(licenseFileNames, "|")))

	commentSyntaxesRe = map[string]*regexp.Regexp {
		"ANTLR": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"C": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"C++": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"C#": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"CSS": regexp.MustCompile(`\/\*(.*?\t?\r?\n?)+?\*\/`),
		"Go": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"HTML": regexp.MustCompile(`<\!--(.*?\t?\r?\n?)+?-->`),
		"Haskel": regexp.MustCompile(`(-{2}.*\t?\r?\n?)|(\{-(.*?\t?\r?\n?)+?\-\})`),
		"Java": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"JavaScript": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"Matlab": regexp.MustCompile(`(%.*\t?\r?\n?)|(%\{(.?\t?\r?\n?)+?%\})`),
		"Objective-C": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"Perl": regexp.MustCompile(`(#.*\t?\r?\n?)|(=begin(.*?\t?\r?\n?)+?=cut)`),
		"PHP": regexp.MustCompile(`(#.*\t?\r?\n?)|(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"Python": regexp.MustCompile("('''(.?\t?\r?\n?)+?''')|(#.*\t?\r?\n?)|(\"\"\"(.?\t?\r?\n?)+?\"\"\")"),
		"Ruby": regexp.MustCompile(`(#.*\t?\r?\n?)|(=begin(.*?\t?\r?\n?)+?=end)`),
		"Rust": regexp.MustCompile(`\/\*(.*?\t?\r?\n?)+?\*\/`),
		"R": regexp.MustCompile(`#.*\t?\r?\n?`),
		"Shell": regexp.MustCompile(`#.*\t?\r?\n?`),
		"Swift": regexp.MustCompile(`(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"SAS": regexp.MustCompile(`(\*(.*?\t?\r?\n?)+?;)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"Scala": regexp.MustCompile(`(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"SQL": regexp.MustCompile(`(-{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"TypeScript": regexp.MustCompile(`(\/{2}.*\t?\r?\n?)|(\/\*(.*?\t?\r?\n?)+?\*\/)`),
		"YAML": regexp.MustCompile(`#.*\t?\r?\n?`),
	}
)

// ExtractLicenseFiles returns the list of possible license texts.
// The file names are matched against the template.
// Reader is used to to read file contents.
func ExtractLicenseFiles(files []string, fs filer.Filer) [][]byte {
	candidates := [][]byte{}
	for _, file := range files {
		if licenseFileRe.MatchString(strings.ToLower(paths.Base(file))) {
			text, err := fs.ReadFile(file)
			if len(text) < 128 {
				// e.g. https://github.com/Unitech/pm2/blob/master/LICENSE
				realText, err := fs.ReadFile(string(bytes.TrimSpace(text)))
				if err == nil {
					file = string(bytes.TrimSpace(text))
					text = realText
				}
			}
			if err == nil {
				if preprocessor, exists := filePreprocessors[paths.Ext(file)]; exists {
					text = preprocessor(text)
				}
				candidates = append(candidates, text)
			}
		}
	}
	return candidates
}

// InvestigateLicenseTexts takes the list of candidate license texts and returns the most probable
// reference licenses matched. Each match has the confidence assigned, from 0 to 1, 1 means 100% confident.
func InvestigateLicenseTexts(texts [][]byte) map[string]float32 {
	maxLicenses := map[string]float32{}
	for _, text := range texts {
		candidates := InvestigateLicenseText(text)
		for name, sim := range candidates {
			maxSim := maxLicenses[name]
			if sim > maxSim {
				maxLicenses[name] = sim
			}
		}
	}
	return maxLicenses
}

// InvestigateLicenseText takes the license text and returns the most probable reference licenses matched.
// Each match has the confidence assigned, from 0 to 1, 1 means 100% confident.
func InvestigateLicenseText(text []byte) map[string]float32 {
	return globalLicenseDatabase().QueryLicenseText(string(text))
}

// ExtractReadmeFiles searches for README files.
// Reader is used to to read file contents.
func ExtractReadmeFiles(files []string, fs filer.Filer) [][]byte {
	candidates := [][]byte{}
	for _, file := range files {
		if readmeFileRe.MatchString(strings.ToLower(file)) {
			text, err := fs.ReadFile(file)
			if err == nil {
				if preprocessor, exists := filePreprocessors[paths.Ext(file)]; exists {
					text = preprocessor(text)
				}
				candidates = append(candidates, text)
			}
		}
	}
	return candidates
}

// InvestigateReadmeTexts scans README files for licensing information and outputs the
// probable names using NER.
func InvestigateReadmeTexts(texts [][]byte, fs filer.Filer) map[string]float32 {
	maxLicenses := map[string]float32{}
	for _, text := range texts {
		candidates := InvestigateReadmeText(text, fs)
		for name, sim := range candidates {
			maxSim := maxLicenses[name]
			if sim > maxSim {
				maxLicenses[name] = sim
			}
		}
	}
	return maxLicenses
}

// InvestigateReadmeText scans the README file for licensing information and outputs probable
// names found with Named Entity Recognition from NLP.
func InvestigateReadmeText(text []byte, fs filer.Filer) map[string]float32 {
	return globalLicenseDatabase().QueryReadmeText(string(text), fs)
}

// IsLicenseDirectory indicates whether the directory is likely to contain licenses.
func IsLicenseDirectory(fileName string) bool {
	return licenseDirectoryRe.MatchString(strings.ToLower(fileName))
}

// ExtractSourceFiles searches for source code files and their returns header comments, when available.
// Enry is used to get possible valuable files.
func ExtractSourceFiles(files []string, fs filer.Filer) [][]byte {
	candidates := [][]byte{}
	langs := []string{}
	var empty []byte
	for _, file := range files {
		text, err := fs.ReadFile(file)
		if err == nil {
			lang := enry.GetLanguage(file, empty)
			langs = append(langs, lang)
			if preprocessor, exists := filePreprocessors[paths.Ext(file)]; exists {
				text = preprocessor(text)
			}
			candidates = append(candidates, text)
		}
	}
	if len(candidates) > 0 {
		candidates = ExtractHeaderComments(candidates, langs)
	}
	return candidates
}

// ExtractHeaderComments searches in source code files for header comments and outputs license text on them them.
func ExtractHeaderComments(candidates [][]byte, langs []string) [][]byte {
	comments := [][]byte{}
	var unsupportedTypes string
	for i, candidate := range candidates {
		candidateLang := langs[i]
		candidateHeader := candidate
		if len(candidateHeader) > 1024 {
			candidateHeader = candidate[:1024]
		}
		if reg, exists := commentSyntaxesRe[candidateLang]; exists {
			if match := reg.FindAllString(string(candidateHeader), -1); match != nil {
				var matchText string
				for _, m := range match {
					matchText += m
				}
				comments = append(comments, []byte(matchText))
			}
		} else {
			match, _ := regexp.Match(candidateLang, []byte(unsupportedTypes))
			if match == false {
				unsupportedTypes += candidateLang + ", "
			}
		}
	}
	if len(unsupportedTypes) > 0 {
		unsupportedTypes = unsupportedTypes[:len(unsupportedTypes)-2]
		fmt.Println("The following file types were not investigated for licenses on the comments:", unsupportedTypes + ". ")
	}
	return comments
}

// InvestigateHeaderComments scans the header comments for licensing information and outputs the
// probable names using NER.
func InvestigateHeaderComments(texts [][]byte, fs filer.Filer) map[string]float32 {
	maxLicenses := map[string]float32{}
	for _, text := range texts {
		candidates := InvestigateHeaderComment(text)
		for name, sim := range candidates {
			maxSim := maxLicenses[name]
			if sim > maxSim {
				maxLicenses[name] = sim
			}
		}
	}
	return maxLicenses
}

// InvestigateHeaderComment scans the header comments for licensing information and outputs probable
// names found with Named Entity Recognition from NLP.
func InvestigateHeaderComment(text []byte) map[string]float32 {
	return globalLicenseDatabase().QuerySourceFile(string(text))
}
