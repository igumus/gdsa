.PHONY: all

all: help

## Build:
tidy: ## Tidy project
	@go mod tidy

build: tidy ## Build project
	@go build .

test-vector: build ## Test transducers
	@go clean -testcache
	@go test -v ./... -race -count=1 -run TestVector

test-list: build ## Test transducers
	@go clean -testcache
	@go test -v ./... -race -count=1 -run TestList

test-all: build ## Test transducers
	@go clean -testcache
	@go test -v ./... -race -count=1

pre-commit: test ## Checks everything is allright
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
