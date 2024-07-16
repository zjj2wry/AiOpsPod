package main

import (
	"github.com/zjj2wry/AiOpsPod/config"
	"github.com/zjj2wry/AiOpsPod/document"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	config, err := config.LoadConfig(".")
	if err != nil {
		logger.Fatal("Error initializing config", zap.Error(err))
	}

	logger.Info("Loaded configuration",
		zap.Any("config", config),
	)

	var lds document.DocumentSource
	if config.Document != nil {
		if config.Document.LocalDir != nil {
			lds = &document.LocalDocumentSource{
				LocalDir: *config.Document.LocalDir,
				Logger:   logger,
			}
		}
	}

	_, err = lds.FetchDocuments()
	if err != nil {
		logger.Error("Error fetching documents", zap.Error(err))
	}
}
