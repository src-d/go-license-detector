package filer

import (
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testFiler(t *testing.T, filer Filer) {
	defer filer.Close()
	files, err := filer.ReadDir("")
	sort.Slice(files, func(i int, j int) bool {
		return files[i].Name < files[j].Name
	})
	assert.Nil(t, err)
	assert.Len(t, files, 2)
	assert.Equal(t, "one", files[0].Name)
	assert.False(t, files[0].IsDir)
	assert.Equal(t, "two", files[1].Name)
	assert.True(t, files[1].IsDir)
	content, err := filer.ReadFile("one")
	assert.Nil(t, err)
	assert.Equal(t, "hello\n", string(content))
	files, err = filer.ReadDir("two")
	assert.Len(t, files, 1)
	assert.Equal(t, "three", files[0].Name)
	assert.False(t, files[0].IsDir)
	content, err = filer.ReadFile("two/three")
	assert.Nil(t, err)
	assert.Equal(t, "world\n", string(content))

	files, err = filer.ReadDir("..")
	assert.Nil(t, files)
	assert.NotNil(t, err)

	files, err = filer.ReadDir("two/three")
	assert.Nil(t, files)
	assert.NotNil(t, err)

	content, err = filer.ReadFile("two/four")
	assert.Nil(t, content)
	assert.NotNil(t, err)
}

func TestLocalFiler(t *testing.T) {
	filer, err := FromDirectory("test_data/local")
	assert.Nil(t, err)
	testFiler(t, filer)
	filer, err = FromDirectory("test_data/local2")
	assert.Nil(t, filer)
	assert.NotNil(t, err)
	filer, err = FromDirectory("test_data/local/one")
	assert.Nil(t, filer)
	assert.NotNil(t, err)
}

func TestGitFiler(t *testing.T) {
	filer, err := FromGitURL("test_data/git")
	assert.Nil(t, err)
	testFiler(t, filer)
	filer, err = FromGitURL("test_data/local2.git")
	assert.Nil(t, filer)
	assert.NotNil(t, err)
}

func TestSivaFiler(t *testing.T) {
	filer, err := FromSiva("test_data/334a82b19a7c893d3807ea52ba35ff2170c296cc.siva")
	assert.Nil(t, err)
	testFiler(t, filer)
	filer, err = FromSiva("test_data/local2.siva")
	defer os.Remove("test_data/local2.siva")
	assert.Nil(t, filer)
	assert.NotNil(t, err)
}

func TestZipFiler(t *testing.T) {
	filer, err := FromZIP("test_data/local.zip")
	assert.Nil(t, err)
	testFiler(t, filer)
	filer, err = FromZIP("test_data/local2.zip")
	assert.Nil(t, filer)
	assert.NotNil(t, err)
}

func TestNestedFiler(t *testing.T) {
	filer, err := FromDirectory("test_data/local")
	assert.Nil(t, err)
	filer2 := NestFiler(filer, "two")
	defer filer2.Close()
	files, err := filer2.ReadDir("")
	assert.Nil(t, err)
	assert.Len(t, files, 1)
	assert.Equal(t, "three", files[0].Name)
	assert.False(t, files[0].IsDir)
	content, err := filer2.ReadFile("three")
	assert.Nil(t, err)
	assert.Equal(t, "world\n", string(content))
}
