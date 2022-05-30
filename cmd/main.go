package main

import (
	"context"
	"go.uber.org/zap"
	"sample-log/pkg/logger"
	"time"
)

func main() {
	ctx := context.Background()

	sugarLogger := logger.ConfigZap()
	defer sugarLogger.Sync()

	for i := 0; i < 10000; i++ {
		if i%3 == 0 {
			logger.Info(ctx, "Info log test", zap.Int("i", i))
		} else if i%3 == 1 {
			logger.Error(ctx, "Error log test", nil, zap.Int("i", i))
		} else {
			logger.Warn(ctx, "Warn log test", zap.Int("i", i))
		}
		time.Sleep(200 * time.Millisecond)
	}
}
