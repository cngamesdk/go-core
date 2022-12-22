package log

import (
	"go.uber.org/zap"
	"testing"
)

func TestFileLog(t *testing.T) {
	log := CreateFileLoggerByLevel(&FileOption{FileName: "/your/app/test.log"})
	log.Info("Come", zap.String("request_id", "123"))
}
