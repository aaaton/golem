SHELL:=/bin/bash
default: all
LANG=en
all:
	# go get -u github.com/jteeuwen/go-bindata/...
	mkdir -p data
	$(MAKE) en
	$(MAKE) sv
	$(MAKE) fr
	$(MAKE) es
	$(MAKE) de
	$(MAKE) it
	

en: LANG=en
en: download
en: package

sv: LANG=sv
sv: download
sv: package

fr: LANG=fr
fr: download
fr: package

es: LANG=es
es: download
es: package

de: LANG=de
de: download
de: package

it: LANG=it
it: download
it: package

download:
	curl https://raw.githubusercontent.com/michmech/lemmatization-lists/master/lemmatization-$(LANG).txt > data/$(LANG)

package:
	go run dicts/cmd/cedar/cedar_creator.go data/$(LANG) data/$(LANG).gob
	go-bindata -o dicts/$(LANG)/$(LANG).go -pkg $(LANG) data/$(LANG).gob
	go run dicts/cmd/generate_pack.go -locale $(LANG) > dicts/$(LANG)/pack.go

benchcmp:
	# ensure no govenor weirdness
	# sudo cpufreq-set -g performance
	go test -test.benchmem=true -run=NONE -bench=. ./... > bench_current.test
	git stash save "stashing for benchcmp"
	@go test -test.benchmem=true -run=NONE -bench=. ./... > bench_head.test
	git stash pop
	benchcmp bench_head.test bench_current.test

profile:
	@mkdir -p pprof/
	go test -run=NONE -cpuprofile pprof/cpu.prof -memprofile pprof/mem.prof -bench .
	go tool pprof -pdf pprof/cpu.prof > pprof/cpu.pdf
	xdg-open pprof/cpu.pdf
	go tool pprof -weblist=.* pprof/cpu.prof
