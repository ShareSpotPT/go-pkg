package log

import (
	"github.com/ShareSpotPT/go-pkg/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	GetGormLogger() GormLogger
}

func ProvideLogger(mode env.Mode) (Logger, error) {
	timeLayout := zapcore.TimeEncoderOfLayout("Mon 02 Jan 15:04:05 MST")

	switch mode {
	case env.Dev:
		logger := &zapLogger{}
		var err error

		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeTime = timeLayout
		config.DisableStacktrace = true
		logger.Logger, err = config.Build(zap.AddCaller())
		if err != nil {
			return logger, err
		}

		return logger, nil

	case env.CICD:
		logger := &zapLogger{}
		var err error

		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.EncoderConfig.EncodeTime = timeLayout
		config.DisableStacktrace = true
		logger.Logger, err = config.Build(zap.AddCaller())
		if err != nil {
			return logger, err
		}

		return logger, nil

	case env.Staging:
		logger := &zapLogger{}
		var err error

		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = timeLayout
		config.DisableStacktrace = true
		logger.Logger, err = config.Build()
		if err != nil {
			return logger, err
		}

		return logger, nil

	case env.Main:
		logger := &zapLogger{}
		var err error

		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = timeLayout
		config.DisableStacktrace = true
		logger.Logger, err = config.Build()
		if err != nil {
			return logger, err
		}

		return logger, nil

	default:
		return nil, env.ErrBadMode("cannot create logger, mode is invalid: " + mode.String())
	}
}

func MockLogger() Logger {
	return &zapLogger{zap.NewNop()}
}
