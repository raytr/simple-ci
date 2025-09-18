
# Go parameters
GOBUILD=go build
GOCLEAN=go clean
GOTEST=go test
GOGET=go get
GOMOD=go mod
BINARY_NAME=simple-ci
BINARY_UNIX=$(BINARY_NAME)_unix

# Coverage settings
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html
COVERAGE_THRESHOLD=80

# Linting tools
GOLANGCI_LINT=golangci-lint
GOSEC=gosec

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...

# Build for Linux
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

# Download dependencies
deps:
	$(GOMOD) download

# Verify dependencies
verify:
	$(GOMOD) verify

# Run tests
test:
	$(GOTEST) -race ./...

# Generate HTML coverage report
coverage-html: coverage
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "HTML coverage report generated: $(COVERAGE_HTML)"

# Format code
format:
	@echo "🔧 Formatting code..."
	gofmt -s -w .
	@echo "✅ Code formatted"

# Check if code is formatted
format-check:
	@echo "🔍 Checking code formatting..."
	@if [ "$$(gofmt -s -l . | wc -l)" -gt 0 ]; then \
		echo "❌ The following files are not formatted correctly:"; \
		gofmt -s -l .; \
		echo "Run 'make format' to fix formatting"; \
		exit 1; \
	else \
		echo "✅ All files are properly formatted"; \
	fi

# Run go vet
vet:
	@echo "🔍 Running go vet..."
	go vet ./...
	@echo "✅ go vet passed"


# Complete lint flow
lint: format-check vet
	@echo "✅ All linting checks passed!"


# Run CI-like checks locally
ci: clean deps verify lint coverage integration
	@echo "🎉 All CI checks passed! Ready to push."

# Development workflow
dev: deps format lint test
	@echo "🚀 Development checks completed!"

