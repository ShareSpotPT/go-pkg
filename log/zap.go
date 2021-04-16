package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm/logger"
	"time"
)

type zapLogger struct {
	*zap.Logger
}

type GormLogger struct {
	*zapLogger
}

func (s *zapLogger) GetGormLogger() GormLogger {
	return GormLogger{s}
}

func (g GormLogger) Trace(ctx  context.Context, begin time.Time, fc func() (string, int64), err error) {
	s, i := fc()
	g.zapLogger.Debug(s, Any("count", i))
}

func (g GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return g
}

func (g GormLogger) Info(ctx  context.Context, s string, i ...interface{}) {
	g.zapLogger.Info(s, zap.Any("", i))
}

func (g GormLogger) Warn(ctx  context.Context, s string, i ...interface{}) {
	g.zapLogger.Warn(s, zap.Any("", i))
}

func (g GormLogger) Error(ctx context.Context, s string, i ...interface{}) {
	g.zapLogger.Error(s, zap.Any("", i))
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

func (s *zapLogger) GetGinLog() *zap.Logger {
	return s.Logger.With(zap.String("router", "gin"))
}