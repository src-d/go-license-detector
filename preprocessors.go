package licenseng

import (
	"bytes"

	"github.com/hhatto/gorst"
	"gopkg.in/russross/blackfriday.v2"
)

func PreprocessMarkdown(text string) string {
	html := blackfriday.Run([]byte(text))
	return HTML2Text(string(html))
}

func PreprocessRestructuredText(text string) string {
	parser := rst.NewParser(nil)
	input := bytes.NewBufferString(text)
	output := &bytes.Buffer{}
	parser.ReStructuredText(input, rst.ToHTML(output))
	return HTML2Text(output.String())
}
