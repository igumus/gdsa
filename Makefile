.PHONY: all

all: help

## Build:
build: ## Builds project
	@go mod tidy
	@go build .

test-clean: build 
	@go clean -testcache

test-all: test-clean ## Runs all tests 
	@go test -v ./... -race -count=1

pre-commit: test-all ## Checks everything is allright
	@echo "Commit Status: OK"

test-transducer: test-clean ## Runs transducer tests
	@go test -v ./... -race -count=1 -run TestFunction

test-iter: test-clean ## Runs iterator tests
	@go test -v ./... -race -count=1 -run TestIterator

test-list: test-clean ## Runs list collection tests 
	@go test -v ./... -race -count=1 -run TestList

test-vector: test-clean ## Runs vector2d tests
	@go test -v ./... -race -count=1 -run TestVector

## Help:
help: ## Show this help.
	@echo ''
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    %-20s%s\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  %s\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
