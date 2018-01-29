# licenseng

Project license detector. It scans the given directory for license files, reads them and outputs
all findings in the standardized way ([SPDX](https://spdx.org/licenses/)). Acts as either
a command line application or a library.

Why? There are no similar projects exist which can be compiled into a native binary without
dependencies. The implementation is also fast and sophisticated.

The license texts are taken directly from [license-list-data](https://github.com/spdx/license-list-data)
SPDX project, thus ≈400 licenses are handled. The detection is fuzzy and is not template matching;
in other words, this project does not provide any legal guarantees. The intended area of usage is
data mining.

## Installation

```
go get -v -d github.com/vmarkovtsev/licenseng/...
cd $GOPATH/src/github.com/vmarkovtsev/licenseng
make dependencies
go get -v github.com/vmarkovtsev/licenseng/...
```

## Contributions

...are welcome, see [CONTRIBUTING.md](CONTRIBUTING.md) and [code of conduct](CODE_OF_CONDUCT.md).

## License

Apache 2.0, see [LICENSE.md](LICENSE.md).

## Algorithm

1. Find files in the root directory which may represent a license. E.g. `LICENSE` or `license.md`.
2. If the file is Markdown or reStructuredText, render to HTML and then convert to plain text.
3. Normalize the text according to [SPDX recommendations](https://spdx.org/spdx-license-list/matching-guidelines).
4. Split the text into unigrams and build the weighted bag of words.
5. Calculate [Weighted MinHash](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/36928.pdf).
6. Apply Locality Sensitive Hashing and pick the reference licenses which are close.
7. For each of the candidate, calculate the [Levenshtein distance](https://en.wikipedia.org/wiki/Levenshtein_distance) - `D`.
the corresponding text is the single line with each unigram represented by a single rune (character).
8. Set the similarity as `1 - D / L` where `L` is the number of unigrams in the quieried license.

This pipeline guarantees constant time queries, though requires some initialization to preprocess
the reference licenses.

## Usage

Command line:

```bash
licenseng /path/to/project
```

Library:

```go
import "github.com/vmarkovtsev/licenseng"

func main() {
	licenses, err := licenseng.InvestigateProjectLicenses("/path/to/project")
}
```