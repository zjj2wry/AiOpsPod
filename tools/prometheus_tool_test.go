package tools

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestPrometheusTool_QueryRange(t *testing.T) {
	// Start the mock Prometheus server
	server, err := StartMockPrometheusServer() // Let it choose a random available port
	if err != nil {
		t.Fatalf("Failed to start mock Prometheus server: %v", err)
	}
	defer server.Close()

	// Create logger
	logger, _ := zap.NewDevelopment()

	// Create PrometheusTool instance
	tool, err := NewPrometheusTool(server.URL, logger, WithTimeout(10*time.Second))
	if err != nil {
		t.Fatalf("Failed to create PrometheusTool: %v", err)
	}

	assert.Equal(t, "10s", tool.options.timeout.String())

	// Define time range
	startTime := time.Unix(1609459200, 0) // Example start time
	endTime := time.Unix(1609459260, 0)   // Example end time
	step := 60 * time.Second

	// Execute range query
	ctx := context.Background()
	output, err := tool.queryRange(ctx, "up", startTime, endTime, step)
	if err != nil {
		t.Fatalf("Error executing query range: %v", err)
	}

	expectedOutput := `up{instance="localhost:9090", job="prometheus"} =>
1 @[1609459200]
1 @[1609459260]`
	assert.Equal(t, expectedOutput, output)
}
