package main

import (
	"github.com/zjj2wry/AiOpsPod/config"

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
}
