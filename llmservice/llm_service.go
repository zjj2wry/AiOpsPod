package llmservice

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/zjj2wry/AiOpsPod/document"
)

type LLMService struct {
	vectorstores.VectorStore
	llms.Model
}

func (d *LLMService) UpdateDocuments(docs []document.Document) {
}

func (d *LLMService) getSimilarDocument(content string) (*document.Document, error) {
	return &document.Document{}, nil
}

func (d *LLMService) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	doc, err := d.getSimilarDocument(prompt)
	if err != nil {
		return "", err
	}
	newPromt := fmt.Sprintf("Context: \n%s\n+%s", doc.Content, prompt)
	return d.Model.Call(ctx, newPromt, options...)
}
