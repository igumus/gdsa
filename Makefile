.PHONY: all

all: help

## Build:
build: tidy ## Build project
	@go mod tidy
	@go build .

test-vector: build ## Runs vector2d tests
	@go clean -testcache
	@go test -v ./... -race -count=1 -run TestVector

test-list: build ## Runs list collection tests 
	@go clean -testcache
	@go test -v ./... -race -count=1 -run TestList

test-iter: build ## Runs iterator tests
	@go clean -testcache
	@go test -v ./... -race -count=1 -run TestIterator

test-all: build ## Runs all tests 
	@go clean -testcache
	@go test -v ./... -race -count=1

pre-commit: test-all ## Checks everything is allright
	@echo "Commit Status: OK"

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
