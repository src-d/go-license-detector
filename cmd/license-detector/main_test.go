package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-license-detector.v3/licensedb"
)

func TestCmdMain(t *testing.T) {
	buffer := &bytes.Buffer{}
	detect([]string{"../..", "."}, "json", buffer)
	var r []licensedb.Result
	json.Unmarshal(buffer.Bytes(), &r)
	assert.Len(t, r, 2)
	assert.Equal(t, "../..", r[0].Arg)
	assert.Equal(t, ".", r[1].Arg)
	assert.Len(t, r[0].Matches, 4)
	assert.Len(t, r[1].Matches, 0)
	assert.Equal(t, "", r[0].ErrStr)
	assert.Equal(t, "no license file was found", r[1].ErrStr)
	assert.Equal(t, "Apache-2.0", r[0].Matches[0].License)
	assert.InDelta(t, 0.9877, r[0].Matches[0].Confidence, 0.001)
	assert.Equal(t, "ECL-2.0", r[0].Matches[1].License)
	assert.InDelta(t, 0.9047, r[0].Matches[1].Confidence, 0.001)
	buffer.Reset()
	detect([]string{"../..", "."}, "text", buffer)
	assert.Equal(t, `../..
	99%	Apache-2.0
	90%	ECL-2.0
	85%	SHL-0.51
	85%	SHL-0.5
.
	no license file was found
`, buffer.String())
}
