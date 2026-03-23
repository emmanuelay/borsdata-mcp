.PHONY: build clean test test-integration install help scan lint mcpb

# Binary name
BINARY_NAME=borsdata-mcp

# Build directory
BUILD_DIR=.

# Source directory
CMD_DIR=.

# Version information
VERSION?=dev
COMMIT?=$(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Build flags
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.date=$(DATE)"

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-18s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the binary
	@echo "Building $(BINARY_NAME)..."
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)
	@echo "Binary created at $(BUILD_DIR)/$(BINARY_NAME)"

clean: ## Remove build artifacts
	@echo "Cleaning..."
	@rm -f $(BUILD_DIR)/$(BINARY_NAME)
	@rm -rf dist/
	@echo "Cleaned"

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-integration: ## Run integration tests (requires network)
	@echo "Running integration tests..."
	@go test -v -tags=integration ./...

install: ## Install the binary to $GOPATH/bin
	@echo "Installing $(BINARY_NAME)..."
	@go install $(LDFLAGS) $(CMD_DIR)
	@echo "Installed to $(GOPATH)/bin/$(BINARY_NAME)"

run: build ## Build and run the MCP server via stdio
	@echo "Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies updated"

scan: ## Run security scan with govulncheck
	@echo "Running vulnerability scan..."
	@govulncheck ./...

lint: ## Run static analysis with staticcheck
	@echo "Running static analysis..."
	@staticcheck ./...

mcpb: ## Build .mcpb bundles from goreleaser dist/
	@scripts/build-mcpb.sh $(VERSION)

all: clean deps build test ## Clean, download deps, build, and test
