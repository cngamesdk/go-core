package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

// FileLog 创建文件日志
func FileLog(fileName string, maxAge time.Duration, rotationTime time.Duration) *logrus.Logger {
	myRotatelogs, rotatelogsErr := rotatelogs.New(
		fileName,
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)

	if rotatelogsErr != nil {
		log.Println(rotatelogsErr.Error())
	}
	logLogrus := logrus.New()
	logLogrus.SetReportCaller(true)
	logLogrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logLogrus.SetOutput(myRotatelogs)
	return logLogrus
}
