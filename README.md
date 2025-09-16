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

## Usage

### Command Line Interface

```bash
# Addition
./simple-ci add 5 3
# Output: Result: 8.00

# Subtraction
./simple-ci subtract 10 4
# Output: Result: 6.00

# Multiplication
./simple-ci multiply 7 6
# Output: Result: 42.00

# Division
./simple-ci divide 15 3
# Output: Result: 5.00

# Division by zero (returns error)
./simple-ci divide 5 0
# Output: Error: division by zero is not allowed
```

### Programmatic Usage

```go
package main

import (
    "fmt"
    "github.com/raytr/simple-ci/calculator"
)

func main() {
    // Basic operations
    result := calculator.Add(5.5, 3.2)
    fmt.Printf("5.5 + 3.2 = %.2f\n", result)
    
    // Division with error handling
    result, err := calculator.Divide(10, 3)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("10 / 3 = %.2f\n", result)
    }
    
    // Advanced operations
    sqrt_result, err := calculator.Sqrt(16)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("√16 = %.2f\n", sqrt_result)
    }
}
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

## Development

### Adding New Features

1. Create a new branch: `git checkout -b feature/new-operation`
2. Add your function to `calculator/calculator.go`
3. Add comprehensive tests to `calculator/calculator_test.go`
4. Update the CLI in `main.go` if needed
5. Run tests locally: `go test ./...`
6. Commit and push your changes
7. Create a pull request

### Code Quality Standards

- All functions must have documentation comments
- Maintain 80%+ test coverage
- Follow Go formatting standards (`gofmt`)
- Pass all linter checks
- Handle errors appropriately
- Include both positive and negative test cases

## Available Operations

### Basic Operations
- `Add(a, b float64) float64` - Addition
- `Subtract(a, b float64) float64` - Subtraction  
- `Multiply(a, b float64) float64` - Multiplication
- `Divide(a, b float64) (float64, error)` - Division with zero check

### Advanced Operations
- `Power(a, b float64) float64` - Exponentiation
- `Sqrt(a float64) (float64, error)` - Square root with negative check
- `Abs(a float64) float64` - Absolute value

### Utility Functions
- `IsEven(n int) bool` - Check if number is even

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## CI/CD Status

The CI pipeline will automatically:
- ✅ Run all unit tests
- ✅ Check code formatting and quality
- ✅ Perform security scans
- ✅ Verify test coverage meets standards
- ✅ Run integration tests
- ❌ Block merge if any checks fail

This ensures that only high-quality, tested code reaches the main branch.
