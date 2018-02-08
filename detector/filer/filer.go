package filer

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// A Filer provides a list of files.
type Filer interface {
	File() (path string, content []byte, err error)
}

type localFiler []string

// FromDirectory returns a Filer that iterates over all the files recursively contained by a directory.
func FromDirectory(path string) (Filer, error) {
	var filer localFiler
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filer = append(filer, path)
		}
		return err
	})
	return &filer, err
}

func (list *localFiler) File() (string, []byte, error) {
	if len(*list) == 0 {
		return "", nil, io.EOF
	}
	path := (*list)[0]
	c, err := ioutil.ReadFile(path)
	*list = (*list)[1:]
	return path, c, err
}

// FromGitURL returns a Filer that iterates over all the files in a git repository given its URL.
func FromGitURL(url string) (Filer, error) {
	repo, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{URL: url})
	if err != nil {
		return nil, errors.Wrapf(err, "could not clone repo from %s", url)
	}
	head, err := repo.Head()
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch HEAD from repo")
	}
	commit, err := repo.CommitObject(head.Hash())
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch commit for HEAD")
	}
	tree, err := commit.Tree()
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch tree for HEAD commit")
	}
	return &gitFiler{tree.Files()}, nil
}

type gitFiler struct{ iter *object.FileIter }

func (f gitFiler) File() (string, []byte, error) {
	next, err := f.iter.Next()
	if err != nil {
		return "", nil, err
	}
	c, err := next.Contents()
	if err != nil {
		return "", nil, err
	}
	return next.Name, []byte(c), nil
}

// WithPathRegexp restricts the files in a given Filer to those with a path matching
// the given regular expression.
func WithPathRegexp(filer Filer, re *regexp.Regexp) Filer { return regexpFiler{filer, re} }

type regexpFiler struct {
	filer Filer
	re    *regexp.Regexp
}

func (f regexpFiler) File() (string, []byte, error) {
	for {
		path, content, err := f.filer.File()
		if err != nil {
			return path, content, err
		}
		if f.re.MatchString(strings.ToLower(path)) {
			return path, content, nil
		}
	}
}

// Contents returns the contents of all of the files in the Filer.
func Contents(filer Filer) ([][]byte, error) {
	var contents [][]byte
	for {
		_, content, err := filer.File()
		if err == io.EOF {
			return contents, nil
		} else if err != nil {
			return nil, err
		}
		contents = append(contents, content)
	}
}
