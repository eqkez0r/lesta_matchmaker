package zaplogger

import (
	"github.com/eqkez0r/lesta_matchmaker/internal/logger"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func New() *ZapLogger {
	return &ZapLogger{
		logger: zap.NewExample().Sugar(),
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

var _ logger.ILogger = (*ZapLogger)(nil)
