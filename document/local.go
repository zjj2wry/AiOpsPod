package document

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/tmc/langchaingo/schema"
	"github.com/zjj2wry/AiOpsPod/config"
	"go.uber.org/zap"
)

var _ DocumentSource = new(LocalDocumentSource)

// LocalDocumentSource implements DocumentSource for local files
type LocalDocumentSource struct {
	config.LocalDir
	Logger *zap.Logger
}

func (lds *LocalDocumentSource) FetchDocuments() ([]schema.Document, error) {
	var documents []schema.Document

	files, err := ioutil.ReadDir(lds.Directory)
	if err != nil {
		return nil, fmt.Errorf("Error reading directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(lds.Directory, file.Name())
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			lds.Logger.Error("Error reading file", zap.Error(err))
			continue
		}

		documents = append(documents, schema.Document{
			Metadata: map[string]interface{}{
				"title": file.Name(),
			},
			PageContent: string(content),
		})
	}

	return documents, nil
}
