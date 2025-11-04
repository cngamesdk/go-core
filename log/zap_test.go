package log

import (
	"github.com/cngamesdk/go-core/config"
	"go.uber.org/zap"
	"testing"
)

func TestZapLog(t *testing.T) {
	config := config.FileLog{}
	config.Filename = "logs/app.log"
	config.MaxSize = 1024
	config.MaxBackups = 30
	config.MaxAge = 90
	config.Compress = true
	logger := NewFileZapLogger(config)
	defer logger.Sync()
	logger.Info("info test", zap.Any("config", config))
	logger.Warn("info warn", zap.Any("config", config))
	logger.Error("info error", zap.Any("config", config))
}
