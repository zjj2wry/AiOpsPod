package document

import (
	"github.com/tmc/langchaingo/schema"
	"go.uber.org/zap"
)

var _ DocumentSource = new(FeishuDocumentSource)

// FeishuDocumentSource implements DocumentSource for Feishu
type FeishuDocumentSource struct {
	Logger *zap.Logger
}

func (fds *FeishuDocumentSource) FetchDocuments() ([]schema.Document, error) {
	return []schema.Document{}, nil
}
