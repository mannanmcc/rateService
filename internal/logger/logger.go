package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var DefaultLogger = zap.New(zapcore.NewCore(
	zapcore.NewJSONEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "@timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
	zapcore.AddSync(os.Stdout),
	zap.NewAtomicLevelAt(zapcore.InfoLevel),
), zap.AddCaller(), zap.AddCallerSkip(1))

// key can be one value only
type key struct{}

// From will return zap looger associated with the context if present, otherwise it will return the default
func From(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(key{}).(*zap.Logger); ok {
		return l
	}

	return DefaultLogger
}

func With(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, key{}, l)
}

func WithFields(ctx context.Context, fields ...zap.Field) context.Context {
	if len(fields) == 0 {
		return ctx
	}

	return With(ctx, From(ctx).With(fields...))
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	From(ctx).Error(msg, fields...)
}

func Print(ctx context.Context, msg string, fields ...zap.Field) {
	From(ctx).Info(msg, fields...)
}

func Sync(ctx context.Context) error {
	return From(ctx).Sync()
}
