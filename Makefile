SHELL:=/bin/bash
default: all
LANG=en
all:
	# go get -u github.com/jteeuwen/go-bindata/...
	mkdir -p data
	$(MAKE) en sv fr es de it

package-all:
	$(MAKE) LANG=en package
	$(MAKE) LANG=sv package
	$(MAKE) LANG=fr package
	$(MAKE) LANG=es package
	$(MAKE) LANG=de package
	$(MAKE) LANG=it package
	
en: 
	$(MAKE) LANG=en download package
sv:  
	$(MAKE) LANG=sv download package
fr:  
	$(MAKE) LANG=fr download package
es:  
	$(MAKE) LANG=es download package
de:  
	$(MAKE) LANG=de download package
it:  
	$(MAKE) LANG=it download package

download:
	curl https://raw.githubusercontent.com/michmech/lemmatization-lists/master/lemmatization-$(LANG).txt > data/$(LANG)

package:
	# Packaging $(LANG)
	go run cmd/simplify/simplify.go data/$(LANG) data/$(LANG).gz
	go run cmd/genpack/genpack.go -locale $(LANG) -path data/$(LANG).gz > v4/dicts/$(LANG)/pack.go
	# ----------------

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
