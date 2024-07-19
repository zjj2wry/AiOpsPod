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

func (d *LLMService) UpdateDocuments(ctx context.Context, docs []schema.Document) {
	d.VectorStore.AddDocuments(ctx, docs)
}

func (d *LLMService) getSimilarDocument(content string) (*schema.Document, error) {
	return &schema.Document{}, nil
}

func (d *LLMService) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	doc, err := d.getSimilarDocument(prompt)
	if err != nil {
		return "", err
	}
	newPromt := fmt.Sprintf("Context: \n%s\n+%s", doc.PageContent, prompt)
	return d.Model.Call(ctx, newPromt, options...)
}
