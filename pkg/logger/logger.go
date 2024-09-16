package logger

import zaplogger "github.com/eqkez0r/lesta_matchmaker/pkg/logger/zap"

type ILogger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Info(v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Error(err error)
}

func New(lType string) ILogger {
	switch lType {
	case "zap":
		{
			return zaplogger.New()
		}
	default:
		{
			return zaplogger.New()
		}
	}
}
