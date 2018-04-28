package fastlog

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastlog(t *testing.T) {
	tests := []float32{
		0.1, 0.5, 0.9, 1.0, 1.1, 2.0, 2.718281828, 3.0, 4.0,
		10.0, 20.0, 100.0, 500.0, 1000.0,
	}
	for _, v := range tests {
		flog := Log(v)
		plog := float32(math.Log(float64(v)))
		if plog != 0 {
			assert.InEpsilon(t, plog, flog, 0.002)
		} else {
			assert.InDelta(t, plog, flog, 0.000002)
		}
	}
}
