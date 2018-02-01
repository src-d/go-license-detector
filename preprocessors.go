package ld

import (
	"bytes"

	"github.com/hhatto/gorst"
	"gopkg.in/russross/blackfriday.v2"
)

// PreprocessMarkdown converts Markdown to plain text. It tries to revert all the decorations.
func PreprocessMarkdown(text string) string {
	html := blackfriday.Run([]byte(text))
	// Repeat to times to heal broken HTML
	return PreprocessHTML(string(html))
}

// PreprocessRestructuredText converts ReStructuredText to plain text.
// It tries to revert all the decorations.
func PreprocessRestructuredText(text string) string {
	parser := rst.NewParser(nil)
	input := bytes.NewBufferString(text)
	output := &bytes.Buffer{}
	parser.ReStructuredText(input, rst.ToHTML(output))
	// Repeat to times to heal broken HTML
	return PreprocessHTML(output.String())
}

// PreprocessHTML converts HTML to plain text. E.g. it rips all the tags.
func PreprocessHTML(text string) string {
	return HTML2Text(HTML2Text(text))
}
