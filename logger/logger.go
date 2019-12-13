package logger

import (
	"fmt"
)
type Zlogger struct {
	log Logger
}

func InitLogger(logway string, filePath string, level string) *Zlogger{
	logger := new(Zlogger)
	switch logway {
	case "FILE":
		logger.log = NewFileLog(filePath, level)
	case "CONSOLE":
		logger.log = NewConsoleLog(level)
	default:
		panic(fmt.Sprintf("unsupport logger name:%s", logway))
	}
	return logger
}

func (z *Zlogger)Debug(format string, args ...interface{}) {
	z.log.Debug(format, args...)
}

func (z *Zlogger)Info(format string, args ...interface{}) {
	z.log.Info(format, args...)
}

func (z *Zlogger)Warn(format string, args ...interface{}) {
	z.log.Warn(format, args...)
}

func (z *Zlogger)Error(format string, args ...interface{}) {
	z.log.Error(format, args...)
}

func (z *Zlogger)Fatal(format string, args ...interface{}) {
	z.log.Fatal(format, args...)
}
