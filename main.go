package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
	"github.com/zjj2wry/AiOpsPod/config"
	"github.com/zjj2wry/AiOpsPod/document"
	"github.com/zjj2wry/AiOpsPod/llmservice"
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

	docs, err := lds.FetchDocuments()
	if err != nil {
		logger.Error("Error fetching documents", zap.Error(err))
	}

	var llm llms.Model
	var vs vectorstores.VectorStore

	if config.LLM.OpenAI != nil {
		openaiClient, err := openai.New(openai.WithToken(config.LLM.OpenAI.Key), openai.WithEmbeddingModel(config.LLM.EmbeddingModel))
		if err != nil {
			logger.Fatal("Error initializing openai client", zap.Error(err))
		}

		llm = openaiClient
		e, err := embeddings.NewEmbedder(openaiClient)
		if err != nil {
			logger.Fatal("Error initializing embedder", zap.Error(err))
		}
		if config.Vector.Weaviate != nil {
			store, err := weaviate.New(
				weaviate.WithScheme(config.Vector.Weaviate.Scheme),
				weaviate.WithHost(config.Vector.Weaviate.Host),
				weaviate.WithEmbedder(e),
				weaviate.WithNameSpace(uuid.New().String()),
				weaviate.WithIndexName("AiOpsPod"+strings.ReplaceAll(uuid.New().String(), "-", "")),
				weaviate.WithQueryAttrs([]string{"title"}))
			if err != nil {
				logger.Fatal("Error initializing weaviate", zap.Error(err))
			}

			vs = store
		}
	}

	llms := llmservice.LLMService{
		VectorStore: vs,
		Model:       llm,
	}

	ctx := context.Background()

	_, err = llms.AddDocuments(ctx, docs)
	if err != nil {
		logger.Fatal("Error add document", zap.Error(err))
	}

	logger.Info("documents added")

	res, err := llms.Call(ctx, "Give a prometheus related sop document")
	if err != nil {
		logger.Fatal("Error answer ops question", zap.Error(err))
	}

	fmt.Println(res)
}
