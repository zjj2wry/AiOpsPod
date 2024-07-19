package llmservice

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
)

type LLMService struct {
	vectorstores.VectorStore
	llms.Model
}

func (d *LLMService) AddDocuments(ctx context.Context, docs []schema.Document) ([]string, error) {
	return d.VectorStore.AddDocuments(ctx, docs)
}

func (d *LLMService) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	docs, err := d.VectorStore.SimilaritySearch(ctx, prompt, 1)
	if err != nil {
		return "", fmt.Errorf("Failed search: %v", err)
	}

	newPromt := fmt.Sprintf("Context: \n%s\n+%s", docs[0].PageContent, prompt)
	return d.Model.Call(ctx, newPromt, options...)
}
