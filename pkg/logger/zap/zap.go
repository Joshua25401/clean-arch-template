package pkg

import (
	"context"

	"go.uber.org/zap"
)

type ZapLogger struct {
	Logger *zap.Logger
}

func (log *ZapLogger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	log.Logger.Debug(msg, fields...)
}

func (log *ZapLogger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	log.Logger.Info(msg, fields...)
}

func (log *ZapLogger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	log.Logger.Warn(msg, fields...)
}

func (log *ZapLogger) Error(ctx context.Context, msg string, err error, fields ...zap.Field) {
	fields = append(fields, zap.Error(err))
	log.Logger.Error(msg, fields...)
}
