package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	*zap.Logger
}

type Field = zapcore.Field


func String(key, val string) Field {
	return zap.String(key, val)
}

func Error(val error) Field {
	return zap.Error(val)
}

func Any(key string, val interface{}) Field {
	return zap.Any(key, val)
}