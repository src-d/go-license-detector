package processors

import (
	"bytes"
	stdhtml "html"
	"strings"

	"golang.org/x/net/html"
)

var (
	skipped = map[string]bool{
		"head":   true,
		"script": true,
		"style":  true,
		"object": true,
	}
	regexBytes = []byte("#$%*/\\\\|><~`=!?.,:;\"'\\])}-")
)

func isHeader(b []byte) bool {
	if len(b) < 2 {
		return false
	}
	return b[0] == 'h' && b[1] >= '2' && b[1] <= '6'
}

func isRegex(b byte) bool { return bytes.ContainsRune(regexBytes, rune(b)) }

// HTML converts HTML to plain text. E.g. it rips all the tags.
func HTML(src []byte) []byte {
	buf := new(bytes.Buffer)

	doc := html.NewTokenizer(bytes.NewReader(src))
	skip := false
	for token := doc.Next(); token != html.ErrorToken; token = doc.Next() {
		name, _ := doc.TagName()
		if skipped[string(name)] {
			if doc.Token().Type != html.SelfClosingTagToken {
				skip = !skip
			}
			continue
		}
		if skip {
			continue
		}

		text := string(doc.Text())
		text = stdhtml.UnescapeString(text)
		text = strings.Replace(text, "\u00a0", " ", -1)
		buf.WriteString(text)

		if string(name) == "br" {
			buf.WriteRune('\n')
			continue
		}
	}

	return buf.Bytes()
}
