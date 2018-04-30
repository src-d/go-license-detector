package wmh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWMHSerialize(t *testing.T) {
	hasher := NewWeightedMinHasher(100, 50, 7)
	bytes, err := hasher.MarshalBinary()
	assert.Nil(t, err)
	newHasher := &WeightedMinHasher{}
	err = newHasher.UnmarshalBinary(bytes)
	assert.Nil(t, err)
	assert.Equal(t, hasher.Bitness, newHasher.Bitness)
	assert.Equal(t, hasher.dim, newHasher.dim)
	assert.Equal(t, hasher.sampleSize, newHasher.sampleSize)
	assert.Equal(t, hasher.rs, newHasher.rs)
	assert.Equal(t, hasher.lnCs, newHasher.lnCs)
	assert.Equal(t, hasher.betas, newHasher.betas)
}

func TestWMHHash(t *testing.T) {
	hasher := NewWeightedMinHasher(100, 50, 7)
	assert.NotNil(t, hasher)
	hasher.Bitness = 32
	hash := hasher.Hash([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90})
	/*
		import numpy, datasketch
		gen = datasketch.WeightedMinHashGenerator(100, 50, 7)
		with open("test_data/wmh.bin", "rb") as fin:
			fin.read(9)
			gen.rs = numpy.frombuffer(fin.read(100*50*4), dtype=numpy.float32).reshape(50, 100)
			gen.ln_cs = numpy.frombuffer(fin.read(100*50*4), dtype=numpy.float32).reshape(50, 100)
			betas = numpy.frombuffer(fin.read(100*50*2), dtype=numpy.uint16)
			gen.betas = (betas / ((1 << 16) - 1)).astype(numpy.float32).reshape(50, 100)
		v = numpy.zeros(100, numpy.float32)
		for i, ii in enumerate([0, 10, 20, 30, 40, 50, 60, 70, 80, 90]):
			v[ii] = i + 1
		mh = gen.minhash(v)
		for h in mh.hashvalues:
			print("%d," % (h[0] | (h[1] << 16)))
	*/
	truth := []uint64{
		65586,
		0,
		65626,
		65616,
		65626,
		30,
		65616,
		90,
		40,
		65576,
		65596,
		65586,
		65626,
		65626,
		589884,
		20,
		65616,
		65626,
		65596,
		65626,
		262234,
		131152,
		65596,
		65596,
		65556,
		65626,
		65576,
		65606,
		65626,
		65606,
		10,
		90,
		65596,
		65586,
		65626,
		65606,
		65626,
		0,
		131162,
		65626,
		65576,
		65626,
		65616,
		65606,
		65606,
		131152,
		65566,
		65626,
		65586,
		65626,
	}
	assert.Equal(t, truth, hash)
}

func TestWMHTrash(t *testing.T) {
	hasher := NewWeightedMinHasher(100, 50, 7)
	assert.Panics(t, func() {
		hasher.Hash([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90})
	})
	assert.Panics(t, func() {
		hasher.Hash([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 10, 20, 30, 40, 50, 60, 70, 80})
	})
	assert.Panics(t, func() {
		hasher.Hash([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 10, 20, 30, 40, 50, 60, 70, 80, 100})
	})
}
