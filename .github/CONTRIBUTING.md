# Contributing to quickwit-go

Thank you for your interest in contributing to quickwit-go! This document provides guidelines and instructions for contributing.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/quickwit-go.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Run tests: `go test -v ./...`
6. Run linter: `golangci-lint run`
7. Commit your changes: `git commit -am 'Add some feature'`
8. Push to the branch: `git push origin feature/your-feature-name`
9. Create a Pull Request

## Development Setup

### Prerequisites

- Go 1.24 or later
- Docker (for running integration tests)
- golangci-lint (for linting)

### Install Dependencies

```bash
go mod download
```

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run tests with race detection
go test -v -race ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...

# Run tests with container logs
CONTAINER_LOG=true go test -v ./...
```

### Running Linter

```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run

# Run linter with auto-fix
golangci-lint run --fix
```

## Code Guidelines

### Go Style

- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` to format your code
- Add comments to exported functions, types, and constants
- Keep functions small and focused
- Use meaningful variable and function names

### Testing

- Write tests for all new features
- Maintain or improve code coverage
- Use table-driven tests where appropriate
- Use Testcontainers for integration tests
- Clean up resources in tests (use `defer` or `t.Cleanup()`)

### Commits

- Write clear, concise commit messages
- Use conventional commit format:
  - `feat: add new feature`
  - `fix: resolve bug`
  - `docs: update documentation`
  - `test: add tests`
  - `refactor: refactor code`
  - `chore: update dependencies`

### Pull Requests

- Fill out the Pull Request template
- Ensure all CI checks pass
- Request review from maintainers
- Address review comments promptly
- Keep PRs focused and reasonably sized

## Project Structure

```
quickwit-go/
├── .github/              # GitHub configuration
│   ├── workflows/        # GitHub Actions workflows
│   ├── CI.md            # CI documentation
│   ├── CONTRIBUTING.md  # This file
│   ├── PULL_REQUEST_TEMPLATE.md
│   └── dependabot.yml   # Dependabot configuration
├── model_*.go           # Model definitions (one per file)
├── client.go            # Client implementation
├── client_test.go       # Client tests
├── builder.go           # Client builder
├── constants.go         # Constants
├── testcontainer.go     # Testcontainer setup
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── README.md            # Project README
├── TESTING.md           # Testing documentation
├── LICENSE              # MIT License
└── .golangci.yml        # Linter configuration
```

## Adding New Features

When adding new features:

1. **Add the method to the `Client` interface** in `client.go`
2. **Implement the method** in `client.go`
3. **Add models** if needed (create new `model_*.go` files)
4. **Write tests** in `client_test.go`
5. **Update documentation** in README.md and TESTING.md
6. **Run tests and linter** to ensure everything works

### Example: Adding a New Endpoint

```go
// 1. Add to Client interface
type Client interface {
    // ... existing methods
    NewMethod(ctx context.Context, param string) (*Response, error)
}

// 2. Implement the method
func (c *client) NewMethod(ctx context.Context, param string) (*Response, error) {
    req, err := http.NewRequestWithContext(
        ctx,
        "GET",
        fmt.Sprintf("%s/api/v1/new-endpoint?param=%s", c.endpoint, param),
        nil,
    )
    if err != nil {
        return nil, err
    }

    for _, interceptor := range c.interceptors {
        interceptor(req)
    }

    return Request[Response](c.log, req)
}

// 3. Add tests
func TestClient(t *testing.T) {
    // ... setup

    t.Run("New Method", func(t *testing.T) {
        result, err := client.NewMethod(ctx, "test-param")
        require.NoError(t, err)
        assert.NotNil(t, result)
    })
}
```

## Reporting Issues

- Use the GitHub issue tracker
- Check if the issue already exists
- Provide a clear description
- Include steps to reproduce
- Include Go version and OS information
- Include relevant logs and error messages

## Questions?

- Open a GitHub issue with the "question" label
- Check existing issues and documentation first

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
