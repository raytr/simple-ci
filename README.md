# Simple CI - Go Calculator with GitHub Actions

A simple calculator application built in Go with comprehensive unit tests and CI/CD pipeline using GitHub Actions.

## Features

- ✅ Basic mathematical operations (Add, Subtract, Multiply, Divide)
- ✅ Advanced operations (Power, Square Root, Absolute Value)
- ✅ Utility functions (Even number checker)
- ✅ Comprehensive error handling
- ✅ Command-line interface
- ✅ 100% test coverage
- ✅ CI/CD pipeline with GitHub Actions
- ✅ Security scanning
- ✅ Code quality checks

## Project Structure

```
simple-ci/
├── calculator/
│   ├── calculator.go      # Core calculator logic
│   └── calculator_test.go # Comprehensive unit tests
├── .github/
│   └── workflows/
│       └── ci.yml         # GitHub Actions CI pipeline
├── main.go               # CLI application
├── go.mod               # Go module definition
└── README.md           # This file
```

## Installation

1. Clone the repository:
```bash
git clone https://github.com/raytr/simple-ci.git
cd simple-ci
```

2. Download dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o simple-ci .
```

## Testing

### Run Unit Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Run Benchmarks

```bash
go test -bench=. ./calculator
```

### Integration Tests

The CI pipeline includes integration tests that verify the CLI functionality:

```bash
# Build and test CLI
go build -o simple-ci .
./simple-ci add 5 3
./simple-ci divide 10 2
```

## CI/CD Pipeline

### GitHub Actions Workflow

The project includes a comprehensive CI/CD pipeline that runs on:
- Pull requests to the `main` branch
- Pushes to the `main` branch

### Pipeline Features

1. **Multi-version Testing**: Tests against Go 1.21.x and 1.22.x
2. **Dependency Caching**: Speeds up builds by caching Go modules
3. **Code Quality Checks**:
   - `go vet` for static analysis
   - `gofmt` for code formatting
   - `golangci-lint` for comprehensive linting
4. **Security Scanning**: Uses Gosec to scan for security vulnerabilities
5. **Test Coverage**: Enforces 80% minimum test coverage
6. **Integration Testing**: Tests the built CLI application
7. **Coverage Reporting**: Uploads coverage reports to Codecov

### Workflow Jobs

- **test**: Runs unit tests, coverage checks, and integration tests
- **security**: Performs security vulnerability scanning
- **lint**: Runs code quality and formatting checks

### Branch Protection

To enable the CI pipeline to block merges with failing tests:

1. Go to your repository settings
2. Navigate to "Branches" 
3. Add a branch protection rule for `main`
4. Enable "Require status checks to pass before merging"
5. Select the required checks: `test`, `security`, `lint`


## CI/CD Status

The CI pipeline will automatically:
- ✅ Run all unit tests
- ✅ Check code formatting and quality
- ✅ Perform security scans
- ✅ Verify test coverage meets standards
- ✅ Run integration tests
- ❌ Block merge if any checks fail

This ensures that only high-quality, tested code reaches the main branch.
