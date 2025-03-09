package logger

import "log"

type ILogger interface {
	Printf(format string, v ...interface{})
}

type consoleLogger struct {
	logObj *log.Logger
}

func (l *consoleLogger) Printf(format string, v ...interface{}) {
	l.logObj.Printf(format, v...)
}

func InitConsoleLogger() ILogger {
	logObj := log.Default()
	logObj.SetPrefix(":jobAPI: ")
	logObj.SetFlags(log.Lmicroseconds | log.Lmsgprefix)

	return &consoleLogger{
		logObj: logObj,
	}
}
