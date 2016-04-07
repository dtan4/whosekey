build:
	go build

deps:
	go get github.com/Masterminds/glide
	glide install

install:
	go install

PHONY: build
