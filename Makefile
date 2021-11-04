.DEFAULT_GOAL := build

DIRS=bin
BINARY=deluge-remove-after

VERSION=$(shell git describe --tags --always --abbrev=0 --match=v* 2> /dev/null | sed -r "s:^v::g" || echo 0)

$(info $(shell mkdir -p $(DIRS)))
BIN=$(CURDIR)/bin
export GOBIN=$(CURDIR)/bin


help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

fetch: ## Fetches the necessary dependencies to build.
	go mod download
	go mod tidy

clean: ## Cleans up generated files/folders from the build.
	/bin/rm -rfv "dist/" "${BINARY}"

generate-markdown:
	go run *.go --generate-markdown > USAGE.md

build: fetch clean ## Compile binary with static assets embedded.
	go build -ldflags '-d -s -w' -tags netgo -installsuffix netgo -v -o "${BINARY}"

debug: clean
	go run -v *.go --debug
