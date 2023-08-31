package initialize

import (
	"user-api/global"

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

func InitLogger() {
	logger, err := NewLogger()

	if err != nil {
		panic(err)
	}

	// flushes buffer, if any
	defer logger.Sync()

	global.Log = logger.Sugar()
}
