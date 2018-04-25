BINARY := whosekey
BINARY_DIR := bin

build:
	go build -o $(BINARY_DIR)/$(BINARY)

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
	go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep
	dep ensure -v

install:
	go install

PHONY: build deps install
