package licenseng

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeLines(t *testing.T) {
	text := "a\r\nb\rc\n\r"
	assert.Equal(t, "a\nb\nc\n\n", NormalizeLicenseText(text))
}

func TestNormalizeWhitespace(t *testing.T) {
	text := "   a\n b\nc    \n"
	assert.Equal(t, "a\nb\nc\n", NormalizeLicenseText(text))
}

func TestNormalizeQuotesLowerCase(t *testing.T) {
	text := `“You” (or “Your”) shall mean an individual or Legal Entity exercising
permissions granted by this License.`
	assert.Equal(t, `"you" (or "your") shall mean an individual or legal entity exercising
permissions granted by this license.`, NormalizeLicenseText(text))
}

func TestNormalizeLinks(t *testing.T) {
	text := "A <https://fsf.org/> B"
	assert.Equal(t, "a https://fsf.org/ b", NormalizeLicenseText(text))
}
