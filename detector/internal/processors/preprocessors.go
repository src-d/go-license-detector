package processors

import (
	"bytes"

	"github.com/hhatto/gorst"
	"gopkg.in/russross/blackfriday.v2"
)

// Markdown converts Markdown to plain text. It tries to revert all the decorations.
func Markdown(text []byte) []byte {
	html := blackfriday.Run([]byte(text))
	return HTML(html)
}

// RestructuredText converts ReStructuredText to plain text.
// It tries to revert all the decorations.
func RestructuredText(text []byte) []byte {
	parser := rst.NewParser(nil)
	buf := new(bytes.Buffer)
	parser.ReStructuredText(bytes.NewReader(text), rst.ToHTML(buf))
	return HTML(buf.Bytes())
}
