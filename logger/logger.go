package logger

import (
	"fmt"
)

var Log Logger

/*
file, "初始化一个文件日志实例"
console, "初始化console日志实例"
*/
func InitLogger(logway string, filePath string, level string) {
	switch logway {
	case "FILE":
		Log = NewFileLog(filePath, level)
	case "CONSOLE":
		Log = NewConsoleLog(level)
	default:
		panic(fmt.Sprintf("unsupport logger name:%s", logway))
	}
}

func Debug(format string, args ...interface{}) {
	Log.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	Log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	Log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	Log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	Log.Fatal(format, args...)
}
