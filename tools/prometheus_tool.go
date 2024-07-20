package tools

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/tmc/langchaingo/callbacks"
	"github.com/tmc/langchaingo/tools"
)

// PrometheusTool is a tool for querying Prometheus metrics.
type PrometheusTool struct {
	CallbacksHandler callbacks.Handler
	Client           v1.API
	Logger           *zap.Logger
	options          options
}

var _ tools.Tool = PrometheusTool{}

// NewPrometheusTool creates a new instance of PrometheusTool.
func NewPrometheusTool(address string, logger *zap.Logger, opts ...Option) (*PrometheusTool, error) {
	options := &options{}
	for _, opt := range opts {
		opt(options)
	}

	client, err := api.NewClient(api.Config{
		Address: address,
	})
	if err != nil {
		return nil, err
	}

	return &PrometheusTool{
		Client:  v1.NewAPI(client),
		Logger:  logger,
		options: *options,
	}, nil
}

func (t PrometheusTool) Name() string {
	return "Prometheus"
}

func (t PrometheusTool) Description() string {
	return `
	"You have the capability to query Prometheus metrics by PrometheusTool."
	"Action Input required a valid Prometheus query."
	"This tool is only used when querying metrics."`
}

// TODO: make LLM seed input like `prometheus_expr,startTime,endTime,step`
func parseInput(input string) (qeury string, startTime, endTime time.Time, step time.Duration) {
	startTime = time.Now().Add(-10 * time.Minute)
	endTime = time.Now()
	step = 60 * time.Second

	return input, startTime, endTime, step
}

func (t PrometheusTool) Call(ctx context.Context, input string) (string, error) {
	query, start, end, step := parseInput(input)
	return t.queryRange(ctx, query, start, end, step)
}

// QueryRange queries Prometheus for metrics over a specified time range.
func (t PrometheusTool) queryRange(ctx context.Context, query string, startTime, endTime time.Time, step time.Duration) (string, error) {
	if t.CallbacksHandler != nil {
		t.CallbacksHandler.HandleToolStart(ctx, query)
	}

	// Execute Prometheus range query
	result, warnings, err := t.Client.QueryRange(ctx, query, v1.Range{Start: startTime, End: endTime, Step: step}, v1.WithTimeout(t.options.timeout))
	if err != nil {
		if t.CallbacksHandler != nil {
			t.CallbacksHandler.HandleToolError(ctx, err)
		}
		return "", err
	}

	if len(warnings) > 0 {
		t.Logger.Warn("Query warnings", zap.Strings("warnings", warnings))
	}

	if t.CallbacksHandler != nil {
		t.CallbacksHandler.HandleToolEnd(ctx, result.String())
	}

	t.Logger.Info("query range", zap.String("result", result.String()))

	return result.String(), nil
}

// Option is a function that configures the PrometheusTool.
type Option func(*options)

// options is a struct that holds the configuration for PrometheusTool.
type options struct {
	timeout time.Duration
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}
