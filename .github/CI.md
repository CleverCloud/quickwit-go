# Continuous Integration

This repository uses GitHub Actions for continuous integration and testing.

## Workflows

### Test Workflow (`.github/workflows/test.yml`)

Runs on:
- Pull requests to `main` or `master`
- Pushes to `main` or `master`

#### Jobs

**1. Test**
- Runs on: `ubuntu-latest`
- Go version: `1.24`
- Steps:
  - Checkout code
  - Set up Go with caching
  - Download and verify dependencies
  - Run tests with race detection
  - Generate coverage report
  - Upload coverage to Codecov (optional)

**2. Lint**
- Runs on: `ubuntu-latest`
- Go version: `1.24`
- Steps:
  - Checkout code
  - Set up Go with caching
  - Run `golangci-lint` with 5m timeout

**3. Build**
- Runs on: `ubuntu-latest`
- Go version: `1.24`
- Steps:
  - Checkout code
  - Set up Go with caching
  - Build all packages

## Test Environment

### Docker Requirements
The integration tests use [Testcontainers](https://golang.testcontainers.org/) to spin up a Quickwit instance in Docker.

GitHub Actions runners have Docker pre-installed, so tests run automatically in CI.

### Test Configuration
- Timeout: 10 minutes
- Race detection: Enabled
- Coverage: Enabled
- Container logs: Disabled in CI (set `CONTAINER_LOG=false`)

## Linting

The project uses `golangci-lint` with the following enabled linters:
- `gofmt` - Format checking
- `govet` - Go vet
- `errcheck` - Unchecked errors
- `staticcheck` - Static analysis
- `ineffassign` - Ineffectual assignments
- `unused` - Unused code
- `gosimple` - Simplification suggestions
- `misspell` - Spelling mistakes
- `unconvert` - Unnecessary conversions
- `unparam` - Unused function parameters
- `exportloopref` - Loop variable references

Configuration: `.golangci.yml`

## Dependabot

Dependabot is configured to automatically:
- Update Go module dependencies weekly
- Update GitHub Actions versions weekly
- Create PRs with labels: `dependencies`, `go`, `github-actions`

Configuration: `.github/dependabot.yml`

## Status Badges

The following badges are available in the README:

- **Tests**: ![Tests](https://github.com/miton18/quickwit-go/actions/workflows/test.yml/badge.svg)
- **Go Report Card**: [![Go Report Card](https://goreportcard.com/badge/github.com/miton18/quickwit-go)](https://goreportcard.com/report/github.com/miton18/quickwit-go)
- **GoDoc**: [![GoDoc](https://godoc.org/github.com/miton18/quickwit-go?status.svg)](https://godoc.org/github.com/miton18/quickwit-go)

## Local Development

### Running Tests Locally
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

### Running Linter Locally
```bash
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linter
golangci-lint run

# Run linter with auto-fix
golangci-lint run --fix
```

## Troubleshooting

### Docker Issues in CI
If tests fail due to Docker issues in CI:
1. Check GitHub Actions runner status
2. Verify Docker is available in the runner
3. Check Testcontainers logs

### Timeout Issues
If tests timeout in CI:
1. Increase timeout in workflow (currently 10m)
2. Check Quickwit container startup time
3. Verify network connectivity

### Coverage Upload Failures
Coverage upload to Codecov is set to `continue-on-error: true`, so failures won't break the build.

To enable Codecov:
1. Sign up at https://codecov.io
2. Add repository
3. No token needed for public repositories
