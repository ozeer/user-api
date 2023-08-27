package config

import (
	"time"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./api.log",
		"stderr",
		// "stdout",
	}

	return cfg.Build()
}

func Init() {
	logger, err := NewLogger()

	if err != nil {
		panic(err)
	}

	defer logger.Sync() // flushes buffer, if any
	url := "https://imooc.com"

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
