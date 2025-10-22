# quickwit-go

[![Tests](https://github.com/miton18/quickwit-go/actions/workflows/test.yml/badge.svg)](https://github.com/miton18/quickwit-go/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/miton18/quickwit-go)](https://goreportcard.com/report/github.com/miton18/quickwit-go)
[![GoDoc](https://godoc.org/github.com/miton18/quickwit-go?status.svg)](https://godoc.org/github.com/miton18/quickwit-go)

A Go client library for [Quickwit](https://quickwit.io/), a cloud-native search engine.

## Features

- ✅ Full Quickwit API support
- ✅ Index management (create, list, get, delete, clear, describe)
- ✅ Source management (create, delete)
- ✅ Search operations
- ✅ Split operations
- ✅ Cluster health checks
- ✅ Elasticsearch-compatible endpoint
- ✅ Comprehensive test suite with Testcontainers

## Installation

```bash
go get github.com/miton18/quickwit-go
```

## Quick Start

```go
package main

import (
    "context"
    "log"

    quickwit "github.com/miton18/quickwit-go"
)

func main() {
    // Create a new client
    client := quickwit.New(
        quickwit.WithEndpoint("http://localhost:7280"),
    )

    ctx := context.Background()

    // Create an index
    indexConfig := quickwit.IndexConfig{
        Version: "0.9",
        ID:      "my-index",
        DocMapping: quickwit.DocMapping{
            Mode: "dynamic",
            FieldMappings: []quickwit.FieldMapping{
                {
                    Name:    "message",
                    Type:    "text",
                    Indexed: true,
                    Stored:  true,
                },
            },
        },
        IndexingSettings: quickwit.Settings{
            Resources: quickwit.Resources{
                HeapSize: "100MB",
            },
        },
        SearchSettings: quickwit.SearchSettings{
            DefaultSearchFields: []string{"message"},
        },
    }

    index, err := client.CreateIndex(ctx, indexConfig)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Created index: %s", index.Config.ID)

    // Search
    results, err := client.Search(ctx, "my-index", "*")
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Found %d results", results.NumHits)
}
```

## Client Options

```go
// With custom endpoint
client := quickwit.New(
    quickwit.WithEndpoint("http://quickwit.example.com:7280"),
)

// With authentication
client := quickwit.New(
    quickwit.WithEndpoint("http://quickwit.example.com:7280"),
    quickwit.WithBearerToken("your-token"),
)

// With basic auth
client := quickwit.New(
    quickwit.WithEndpoint("http://quickwit.example.com:7280"),
    quickwit.WithBasicAuth("username", "password"),
)

// With custom HTTP client
httpClient := &http.Client{Timeout: 30 * time.Second}
client := quickwit.New(
    quickwit.WithHttpClient(httpClient),
)

// With logger
logger := logrus.New()
client := quickwit.New(
    quickwit.WithLogger(logger),
)
```

## API Coverage

### Cluster Operations
- `GetCluster()` - Get cluster information
- `GetElastic()` - Get Elasticsearch-compatible endpoint info

### Index Operations
- `CreateIndex(ctx, config)` - Create a new index
- `ListIndexes(ctx)` - List all indexes
- `GetIndex(ctx, indexID)` - Get a specific index
- `DeleteIndex(ctx, indexID)` - Delete an index
- `ClearIndex(ctx, indexID)` - Clear all data from an index
- `DescribeIndex(ctx, indexID)` - Get index statistics
- `ListSplits(ctx, indexID)` - List index splits

### Source Operations
- `CreateSource(ctx, indexID, config)` - Create a data source
- `DeleteSource(ctx, indexID, sourceID)` - Delete a source

### Search Operations
- `Search(ctx, indexID, query)` - Execute a search query

## Testing

The library includes comprehensive integration tests using Testcontainers.

```bash
# Run tests
go test -v ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...

# Run tests with container logs
CONTAINER_LOG=true go test -v ./...
```

See [TESTING.md](TESTING.md) for more details.

## Development

### Requirements
- Go 1.24+
- Docker (for integration tests)

### Running Tests Locally
```bash
# Install dependencies
go mod download

# Run all tests
go test -v ./...

# Run linter
golangci-lint run
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Links

- [Quickwit Documentation](https://quickwit.io/docs)
- [Quickwit GitHub](https://github.com/quickwit-oss/quickwit)
