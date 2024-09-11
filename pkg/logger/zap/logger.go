package zaplogger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func New() *ZapLogger {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return &ZapLogger{
		logger: logger.Sugar(),
	}
}

func (z *ZapLogger) Info(v ...interface{}) {
	z.logger.Info(v...)
}

func (z *ZapLogger) Error(err error) {
	z.logger.Error(err)
}

func (z *ZapLogger) Debugf(format string, v ...interface{}) {
	z.logger.Debugf(format, v...)
}

func (z *ZapLogger) Infof(format string, v ...interface{}) {
	z.logger.Infof(format, v...)
}

func (z *ZapLogger) Warnf(format string, v ...interface{}) {
	z.logger.Warnf(format, v...)
}

func (z *ZapLogger) Errorf(format string, v ...interface{}) {
	z.logger.Errorf(format, v...)
}
