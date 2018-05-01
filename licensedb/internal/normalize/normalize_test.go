package normalize

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeLines(t *testing.T) {
	tt := []struct {
		name    string
		in, out string
	}{
		{"lines", "a\r\nb\rc\n\r", "a\nb\nc\n"},
		{"whitespace", "   a\n b\nc    \n", "a\nb\nc\n"},
		{"quotes lowercase",
			`“You” (or “Your”) shall mean an individual or Legal Entity exercising
			permissions granted by this License.`,
			`"you" (or "your") shall mean an individual or legal entity exercising
permissions granted by this license.`},
		{"normalize links", "A <https://fsf.org/> B", "a https:/fsf.org/ b"},
		{"license", "license.\n\nlicence\n\n", "license\n\nlicense\n\n"},
		{"punctuation", "a-‒–—―⁓⸺⸻~˗‐‑⁃⁻₋−∼⎯⏤─➖𐆑֊﹘﹣－", "a-"},
		{"bullet", "-\n*\n✱\n﹡\n•\n●\n⚫\n⏺\n🞄\n∙\n⋅\n", ""},
		{"license", "", ""},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.out, LicenseText(tc.in, Enforced))
	}
}
