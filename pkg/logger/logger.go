package pkg

import (
	pkg "clean-arch-template/pkg/logger/zap"
	"context"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	// Logging level
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, err error, fields ...zap.Field)
}

func NewZapLogger() Logger {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)

	loggerIns := zap.NewProductionConfig()
	loggerIns.EncoderConfig.LevelKey = "severity"
	loggerIns.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	loggerIns.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	loggerIns.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerIns.DisableStacktrace = true

	logger, err := loggerIns.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return nil
	}

	return &pkg.ZapLogger{
		Logger: logger,
	}
}
