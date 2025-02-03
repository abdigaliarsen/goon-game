package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"goon-game/internal/wikipedia/config"
	"goon-game/pkg/utils"
)

func New(cfg *config.Config) utils.Logger {
	return initLogger(cfg)
}

func initLogger(cfg *config.Config) utils.Logger {
	zapConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	switch cfg.LogConfig.ENV {
	case "development":
		zapConfig.Development = true
		zapConfig.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	default:
		zapConfig.Development = false
		zapConfig.EncoderConfig = zap.NewProductionEncoderConfig()
	}

	logger := zap.Must(zapConfig.Build())
	sugar := logger.Sugar()

	return sugar
}
