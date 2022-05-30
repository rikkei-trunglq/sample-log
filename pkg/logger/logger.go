package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugaredLogger *zap.SugaredLogger

//ConfigZap configuration zap
func ConfigZap() *zap.SugaredLogger {

	cfg := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		OutputPaths: []string{"stdout"},

		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "timestamp",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,
			EncodeLevel:  customLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
		},
	}

	logger, _ := cfg.Build()
	logger = logger.With(zap.String("service", "write-log"))
	sugaredLogger = logger.Sugar()
	return sugaredLogger
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// Debug Debug level log
func Debug(_ context.Context, msg string, fields ...interface{}) {
	sugaredLogger.Debugw(msg, fields...)
}

// Info Information level log
func Info(_ context.Context, msg string, fields ...interface{}) {
	sugaredLogger.Infow(msg, fields...)
}

// Warn Warning level log
func Warn(_ context.Context, msg string, fields ...interface{}) {
	sugaredLogger.Warnw(msg, fields...)
}

// Error Error level log
func Error(_ context.Context, msg string, err error, fields ...interface{}) {
	if err != nil {
		fields = append(fields, zap.Any("error", err))
	}
	sugaredLogger.Errorw(msg, fields...)
}
