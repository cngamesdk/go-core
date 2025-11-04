package log

import (
	"context"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logRequestIdKey = "request_id"
)

type MyLogger struct {
	*zap.Logger
	CtxRequestIdKey string
}

func (receiver MyLogger) GetRequestId(ctx context.Context) string {
	return cast.ToString(ctx.Value(receiver.CtxRequestIdKey))
}

func (receiver MyLogger) InfoCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	receiver.With(zap.String(logRequestIdKey, receiver.GetRequestId(ctx))).Info(msg, fields...)
	return
}

func (receiver MyLogger) DebugCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	receiver.With(zap.String(logRequestIdKey, receiver.GetRequestId(ctx))).Debug(msg, fields...)
	return
}

func (receiver MyLogger) WarnCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	receiver.With(zap.String(logRequestIdKey, receiver.GetRequestId(ctx))).Warn(msg, fields...)
	return
}

func (receiver MyLogger) ErrorCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	receiver.With(zap.String(logRequestIdKey, receiver.GetRequestId(ctx))).Error(msg, fields...)
	return
}

func (receiver MyLogger) PanicCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	receiver.With(zap.String(logRequestIdKey, receiver.GetRequestId(ctx))).Panic(msg, fields...)
	return
}

func (receiver MyLogger) DPanicCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	receiver.With(zap.String(logRequestIdKey, receiver.GetRequestId(ctx))).DPanic(msg, fields...)
	return
}
