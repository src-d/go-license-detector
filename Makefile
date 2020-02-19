GOPATH ?= $(shell go env GOPATH)
SPDX_DATA_VERSION ?= 3.8

licensedb/internal/assets/bindata.go: licenses.tar urls.csv names.csv $(GOPATH)/bin/go-bindata
	rm -rf license-list-data-$(SPDX_DATA_VERSION)
	rm -f license-list-data.tar.gz
	$(GOPATH)/bin/go-bindata -nometadata -pkg assets -o licensedb/internal/assets/bindata.go licenses.tar urls.csv names.csv
	rm licenses.tar urls.csv names.csv

licenses.tar: license-list-data.tar.gz
	tar -xf license-list-data.tar.gz license-list-data-$(SPDX_DATA_VERSION)/text
	tar -cf licenses.tar -C license-list-data-$(SPDX_DATA_VERSION)/text .

license-list-data-$(SPDX_DATA_VERSION)/json/details: license-list-data.tar.gz
	tar -xf license-list-data.tar.gz license-list-data-$(SPDX_DATA_VERSION)/json/details

urls.csv: license-list-data-$(SPDX_DATA_VERSION)/json/details
	go run licensedb/internal/assets/extract_urls.go license-list-data-$(SPDX_DATA_VERSION)/json/details > urls.csv

names.csv: license-list-data-$(SPDX_DATA_VERSION)/json/details
	go run licensedb/internal/assets/extract_names.go license-list-data-$(SPDX_DATA_VERSION)/json/details > names.csv

license-list-data.tar.gz:
	curl -SLk -o license-list-data.tar.gz https://github.com/spdx/license-list-data/archive/v$(SPDX_DATA_VERSION).tar.gz

$(GOPATH)/bin/go-bindata:
	go get -v github.com/jteeuwen/go-bindata/go-bindata@6025e8de665b31fa74ab1a66f2cddd8c0abf887e
