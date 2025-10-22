package quickwit

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	QuickwitImage    = "quickwit/quickwit:edge"
	QuickwitHttpPort = "7280/tcp"
	QuickwitGrpcPort = "7281/tcp"
)

var ContainerLogs = os.Getenv("CONTAINER_LOG") == "true"

// QuickwitContainer represents a Quickwit testcontainer instance
type QuickwitContainer struct {
	testcontainers.Container
	Endpoint string
}

// SetupQuickwitContainer starts a Quickwit container for testing
func SetupQuickwitContainer(ctx context.Context) (*QuickwitContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        QuickwitImage,
		ExposedPorts: []string{QuickwitGrpcPort, QuickwitHttpPort},
		Cmd:          []string{"run"},
		WaitingFor: wait.ForAll(
			wait.ForExposedPort(),
			wait.ForHTTP("/health/readyz").
				WithPort(QuickwitHttpPort).
				WithStartupTimeout(60*time.Second),
		),
	}

	if ContainerLogs {
		req.LogConsumerCfg = &testcontainers.LogConsumerConfig{
			Consumers: []testcontainers.LogConsumer{
				&testcontainers.StdoutLogConsumer{},
			},
		}
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, "7280")
	if err != nil {
		return nil, fmt.Errorf("failed to get mapped port: %w", err)
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get host: %w", err)
	}

	endpoint := fmt.Sprintf("http://%s:%s", hostIP, mappedPort.Port())

	return &QuickwitContainer{
		Container: container,
		Endpoint:  endpoint,
	}, nil
}
