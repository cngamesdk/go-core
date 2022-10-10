package log

import (
	"testing"
	"time"
)

func TestFileLog(t *testing.T) {
	log := FileLog(
		"/Applications/projects/cngamesdk/go-core/log/test.%Y%m%d%H%M.log",
		time.Hour*24*90,
		time.Minute,
	)
	log.Info(time.Now())
}
