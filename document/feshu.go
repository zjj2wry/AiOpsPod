package document

import (
	"go.uber.org/zap"
)

var _ DocumentSource = new(FeishuDocumentSource)

// FeishuDocumentSource implements DocumentSource for Feishu
type FeishuDocumentSource struct {
	Logger *zap.Logger
}

func (fds *FeishuDocumentSource) FetchDocuments() ([]Document, error) {
	return []Document{}, nil
}
