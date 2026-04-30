.PHONY: help setup install lint lint-fix test test-coverage build clean run hooks hooks-run deps-check fmt vet dev

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

# Binary name
BINARY_NAME=gotemplate
BINARY_PATH=./bin/$(BINARY_NAME)
CMD_PATH=./cmd/gotemplate

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

setup: install hooks ## Setup project (install deps + git hooks)
	@echo "✅ Project ready — run 'make dev' to start"

install: ## Download dependencies
	$(GOMOD) download
	$(GOMOD) tidy

lint: ## Run golangci-lint
	golangci-lint run ./...

lint-fix: ## Run golangci-lint with auto-fix
	golangci-lint run --fix ./...

fmt: ## Format code
	$(GOFMT) ./...

vet: ## Run go vet
	$(GOCMD) vet ./...

test: ## Run tests
	$(GOTEST) ./... -v -race

test-coverage: ## Run tests with coverage
	$(GOTEST) ./... -v -race -coverprofile=coverage.out
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

build: ## Build the binary
	$(GOBUILD) -o $(BINARY_PATH) -v $(CMD_PATH)

run: build ## Build and run
	$(BINARY_PATH)

dev: ## Run with hot reload
	air

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out coverage.html

hooks: ## Install git hooks via lefthook
	$(GOCMD) tool lefthook install

hooks-run: ## Run lefthook manually
	$(GOCMD) tool lefthook run pre-commit

deps-check: ## Check for outdated dependencies
	$(GOMOD) outdated
