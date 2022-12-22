package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path"
	"strings"
)

type FileOption struct {
	FileName   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
	LocalTime  bool
}

func (receiver *FileOption) initConfig() {

	receiver.LocalTime = true

	if receiver.FileName == "" {
		receiver.FileName = "default"
	}

	fileNameWithSuffix := path.Base(receiver.FileName)
	ext := path.Ext(fileNameWithSuffix)
	receiver.FileName = strings.TrimSuffix(receiver.FileName, ext)

	if receiver.MaxSize <= 0 {
		receiver.MaxSize = 5 * 1024
	}
	if receiver.MaxAge <= 0 {
		receiver.MaxAge = 1
	}
	if receiver.MaxBackups <= 0 {
		receiver.MaxBackups = 30
	}
}

// CreateFileLoggerByLevel 创建文件等级日志.将会产生不同等级日志文件
// 如:debug info warn error panic fatal
func CreateFileLoggerByLevel(fileOption *FileOption) *zap.Logger {

	fileOption.initConfig()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	coreDebug := zapcore.NewCore(encoder, zapcore.AddSync(&lumberjack.Logger{
		LocalTime:  fileOption.LocalTime,
		MaxAge:     fileOption.MaxAge,
		MaxSize:    fileOption.MaxSize,
		Filename:   fileOption.FileName + "-debug.log",
		MaxBackups: fileOption.MaxBackups,
		Compress:   fileOption.Compress,
	}), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.DebugLevel
	}))

	coreInfo := zapcore.NewCore(encoder, zapcore.AddSync(&lumberjack.Logger{
		LocalTime:  fileOption.LocalTime,
		MaxAge:     fileOption.MaxAge,
		MaxSize:    fileOption.MaxSize,
		Filename:   fileOption.FileName + "-info.log",
		MaxBackups: fileOption.MaxBackups,
		Compress:   fileOption.Compress,
	}), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel
	}))

	coreWarn := zapcore.NewCore(encoder, zapcore.AddSync(&lumberjack.Logger{
		LocalTime:  fileOption.LocalTime,
		MaxAge:     fileOption.MaxAge,
		MaxSize:    fileOption.MaxSize,
		Filename:   fileOption.FileName + "-warn.log",
		MaxBackups: fileOption.MaxBackups,
		Compress:   fileOption.Compress,
	}), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.WarnLevel
	}))

	coreError := zapcore.NewCore(encoder, zapcore.AddSync(&lumberjack.Logger{
		LocalTime:  fileOption.LocalTime,
		MaxAge:     fileOption.MaxAge,
		MaxSize:    fileOption.MaxSize,
		Filename:   fileOption.FileName + "-error.log",
		MaxBackups: fileOption.MaxBackups,
		Compress:   fileOption.Compress,
	}), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.ErrorLevel
	}))

	corePanic := zapcore.NewCore(encoder, zapcore.AddSync(&lumberjack.Logger{
		LocalTime:  fileOption.LocalTime,
		MaxAge:     fileOption.MaxAge,
		MaxSize:    fileOption.MaxSize,
		Filename:   fileOption.FileName + "-panic.log",
		MaxBackups: fileOption.MaxBackups,
		Compress:   fileOption.Compress,
	}), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.PanicLevel
	}))

	coreFatal := zapcore.NewCore(encoder, zapcore.AddSync(&lumberjack.Logger{
		LocalTime:  fileOption.LocalTime,
		MaxAge:     fileOption.MaxAge,
		MaxSize:    fileOption.MaxSize,
		Filename:   fileOption.FileName + "-fatal.log",
		MaxBackups: fileOption.MaxBackups,
		Compress:   fileOption.Compress,
	}), zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.FatalLevel
	}))

	return zap.New(zapcore.NewTee(coreDebug, coreInfo, coreWarn, coreError, corePanic, coreFatal), zap.AddCaller())
}

// CreateFileLogger 创建文件日志
func CreateFileLogger(fileOption *FileOption) *zap.Logger {

	fileOption.initConfig()

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}

	// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	core := zapcore.NewCore(encoder, zapcore.AddSync(&lumberjack.Logger{
		LocalTime:  fileOption.LocalTime,
		MaxAge:     fileOption.MaxAge,
		MaxSize:    fileOption.MaxSize,
		Filename:   fileOption.FileName + ".log",
		MaxBackups: fileOption.MaxBackups,
		Compress:   fileOption.Compress,
	}), zapcore.DebugLevel)

	return zap.New(core, zap.AddCaller())
}
