package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	buffer := &bytes.Buffer{}
	detect([]string{"../..", "."}, "json", buffer)
	var r []result
	json.Unmarshal(buffer.Bytes(), &r)
	assert.Len(t, r, 2)
	assert.Equal(t, "../..", r[0].Arg)
	assert.Equal(t, ".", r[1].Arg)
	assert.Len(t, r[0].Matches, 2)
	assert.Len(t, r[1].Matches, 0)
	assert.Nil(t, r[0].Err)
	assert.Equal(t, "", r[0].ErrStr)
	assert.Nil(t, r[1].Err) // json discards
	assert.Equal(t, "no license file was found", r[1].ErrStr)
	assert.Equal(t, "Apache-2.0", r[0].Matches[0].License)
	assert.InDelta(t, 0.9846, r[0].Matches[0].Confidence, 0.001)
	assert.Equal(t, "ECL-2.0", r[0].Matches[1].License)
	assert.InDelta(t, 0.8995, r[0].Matches[1].Confidence, 0.001)
	buffer.Reset()
	detect([]string{"../..", "."}, "text", buffer)
	assert.Equal(t, `../..
	98%	Apache-2.0
	90%	ECL-2.0
.
	no license file was found
`, buffer.String())
}
