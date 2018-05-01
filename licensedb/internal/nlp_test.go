package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitLicenseName(t *testing.T) {
	assert.Equal(t, []substring{
		{"gpl", 1},
		{"3", 1},
	}, splitLicenseName("GPLv3"))
	assert.Equal(t, []substring{
		{"gpl", 1},
		{"3", 2},
	}, splitLicenseName("GPL-3-3"))
	assert.Equal(t, []substring{
		{"apache", 1},
		{"2", 1},
		{"0", 1},
	}, splitLicenseName("Apache-2.0"))
	assert.Equal(t, []substring{
		{"bsd", 1},
		{"2", 1},
	}, splitLicenseName("Simplified BSD"))
	assert.Equal(t, []substring{
		{"gpl", 1},
		{"2", 1},
	}, splitLicenseName("GPL-2-deprecated"))
}
