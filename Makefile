SHELL:=/bin/bash
default: download-all
LANG=en
download-all:
	mkdir -p data
	$(MAKE) en
	$(MAKE) sv
	$(MAKE) fr
	$(MAKE) es
	$(MAKE) de
	rm data/*.zip	
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -o data.go -nocompress data/

en: LANG=en
en: download

sv: LANG=sv
sv: download

fr: LANG=fr
fr: download

es: LANG=es
es: download

de: LANG=de
de: download

download:
	curl http://www.lexiconista.com/Datasets/lemmatization-$(LANG).zip > data/$(LANG).zip
	unzip data/$(LANG).zip -d data
	mv data/lemmatization-$(LANG).txt data/$(LANG)
	gzip data/$(LANG)
