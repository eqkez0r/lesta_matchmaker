package logger

type ILogger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Info(v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Error(err error)
}
