# go-license-detector [![GoDoc](https://godoc.org/gopkg.in/src-d/go-license-detector.v2?status.svg)](http://godoc.org/gopkg.in/src-d/go-license-detector.v2) [![Build Status](https://travis-ci.org/src-d/go-license-detector.svg?branch=master)](https://travis-ci.org/src-d/go-license-detector) [![Build status](https://ci.appveyor.com/api/projects/status/bxixhg99ewcgc5i5?svg=true)](https://ci.appveyor.com/project/vmarkovtsev/go-license-detector) [![codecov](https://codecov.io/github/src-d/go-license-detector/coverage.svg)](https://codecov.io/gh/src-d/go-license-detector) [![Go Report Card](https://goreportcard.com/badge/github.com/src-d/go-license-detector)](https://goreportcard.com/report/github.com/src-d/go-license-detector)

Project license detector - a command line application and a library, written in Go.
It scans the given directory for license files, normalizes and hashes them and outputs
all the fuzzy matches with the list of reference texts.
The returned names follow [SPDX](https://spdx.org/licenses/) standard.
Read the [blog post](https://blog.sourced.tech/post/gld/).

Why? There are no similar projects which can be compiled into a native binary without
dependencies and also support the whole SPDX license database (≈400 items).
This implementation is also fast, requires little memory, and the API is easy to use.

The license texts are taken directly from [license-list-data](https://github.com/spdx/license-list-data)
repository. The detection algorithm is **not template matching**;
this directly implies that go-license-detector does not provide any legal guarantees.
The intended area of it's usage is data mining.

## Installation

```
go get -v gopkg.in/src-d/go-license-detector.v2/...
```

## Contributions

...are welcome, see [CONTRIBUTING.md](CONTRIBUTING.md) and [code of conduct](CODE_OF_CONDUCT.md).

## License

Apache 2.0, see [LICENSE.md](LICENSE.md).

## Algorithm

1. Find files in the root directory which may represent a license. E.g. `LICENSE` or `license.md`.
2. If the file is Markdown or reStructuredText, render to HTML and then convert to plain text. Original HTML files are also converted.
3. Normalize the text according to [SPDX recommendations](https://spdx.org/spdx-license-list/matching-guidelines).
4. Split the text into unigrams and build the weighted bag of words.
5. Calculate [Weighted MinHash](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/36928.pdf).
6. Apply Locality Sensitive Hashing and pick the reference licenses which are close.
7. For each of the candidate, calculate the [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance) - `D`.
the corresponding text is the single line with each unigram represented by a single rune (character).
8. Set the similarity as `1 - D / L` where `L` is the number of unigrams in the quieried license.

This pipeline guarantees constant time queries, though requires some initialization to preprocess
the reference licenses.

If there are not license files found:

1. Look for README files.
2. If the file is Markdown or reStructuredText, render to HTML and then convert to plain text. Original HTML files are also converted.
3. Scan for words like "copyright", "license" and "released under". Take the neighborhood.
4. Run Named Entity Recognition (NER) over that surrounding context and extract the possible license name.
5. Match it against the list of license names from SPDX.

## Usage

Command line:

```bash
license-detector /path/to/project
license-detector https://github.com/src-d/go-git
```

Library:

```go
import (
    "gopkg.in/src-d/go-license-detector.v2/licensedb"
    "gopkg.in/src-d/go-license-detector.v2/licensedb/filer"
)

func main() {
	licenses, err := licensedb.Detect(filer.FromDirectory("/path/to/project"))
}
```

## Quality

On the [dataset](dataset.zip) of ~1000 most starred repositories on GitHub as of early February 2018
([list](dataset.projects.gz)), **99%** of the licenses are detected.
The analysis of detection failures is going in [FAILURES.md](FAILURES.md).

Comparison to other projects on that dataset:

|Detector|Detection rate|Time to scan, sec|
|:-------|:----------------------------------------:|:-----------------------------------------|
|[go-license-detector](https://github.com/src-d/go-license-detector)| 99%  (897/902) | 13.5 |
|[benbalter/licensee](https://github.com/benbalter/licensee)| 75%  (673/902) | 111 |
|[google/licenseclassifier](https://github.com/google/licenseclassifier)| 76%  (682/902) | 907 |
|[boyter/lc](https://github.com/boyter/lc)| 88%  (797/902) | 548 |
|[amzn/askalono](https://github.com/amzn/askalono)| 87%  (785/902) | 165 |
|[LiD](https://source.codeaurora.org/external/qostg/lid)| 94%  (847/902) | 3660 |

<details><summary>How this was measured</summary>
<pre><code>$ cd $(go env GOPATH)/src/gopkg.in/src-d/go-license-detector.v2/licensedb
$ mkdir dataset && cd dataset
$ unzip ../dataset.zip
$ # src-d/go-license-detector
$ time license-detector * \
  | grep -Pzo '\n[-0-9a-zA-Z]+\n\tno license' | grep -Pa '\tno ' | wc -l
$ # benbalter/licensee
$ time ls -1 | xargs -n1 -P4 licensee \
  | grep -E "^License: Other" | wc -l
$ # google/licenseclassifier
$ time find -type f -print | xargs -n1 -P4 identify_license \
  | cut -d/ -f2 | sort | uniq | wc -l
$ # boyter/lc
$ time lc . \
  | grep -vE 'NOASSERTION|----|Directory' | cut -d" " -f1 | sort | uniq | wc -l
$ # amzn/askalono
$ echo '#!/bin/sh
result=$(askalono id "$1")
echo "$1
$result"' > ../askalono.wrapper
$ time find -type f -print | xargs -n1 -P4 sh ../askalono.wrapper | grep -Pzo '.*\nLicense: .*\n' askalono.txt | grep -av "License: " | cut -d/ -f 2 | sort | uniq | wc -l
$ # LiD
$ time license-identifier -I dataset -F csv -O lid
$ cat lid_*.csv | cut -d, -f1 | cut -d"'" -f 2 | grep / | cut -d/ -f2 | sort | uniq | wc -l
</code></pre>
</details>

## Regenerate binary data

The SPDX licenses are included into the binary. To update them, run
```
make bindata.go
```

