package llmservice

import (
	"context"
	"fmt"
	"time"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/zjj2wry/AiOpsPod/config"
	aoptools "github.com/zjj2wry/AiOpsPod/tools"
	"go.uber.org/zap"
)

type LLMService struct {
	config.Config
	vectorstores.VectorStore
	llms.Model
	*zap.Logger
}

func (d *LLMService) AddDocuments(ctx context.Context, docs []schema.Document) ([]string, error) {
	return d.VectorStore.AddDocuments(ctx, docs)
}

func (d *LLMService) Ask(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	// best-effort
	docs, err := d.VectorStore.SimilaritySearch(ctx, prompt, 1, vectorstores.WithScoreThreshold(d.Config.Vector.ScoreThreshold))
	if err != nil {
		d.Logger.Error("Failed search", zap.Error(err))
	}

	var tools []tools.Tool
	if d.Config.Prometheus != nil {
		ptl, err := aoptools.NewPrometheusTool(d.Config.Prometheus.Address, d.Logger, aoptools.WithTimeout(10*time.Second))
		if err != nil {
			return "", fmt.Errorf("Failed create prometheus tool")
		}
		tools = append(tools, ptl)
	}

	content := ""
	if len(docs) > 0 {
		content = docs[0].PageContent
	}

	d.Logger.Info("find related document", zap.String("content", content))

	executor := NewAgent(d.Model, tools, content)

	return chains.Run(context.Background(), executor, prompt)
}
