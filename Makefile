BINPATH ?= build

build:
	go build -tags 'production' -o $(BINPATH)/dp-developer-site

debug:
	go build -tags 'debug' -o $(BINPATH)/dp-developer-site
	HUMAN_LOG=1 DEBUG=1 $(BINPATH)/dp-developer-site

.PHONY: build debug
