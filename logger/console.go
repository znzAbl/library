package logger

import (
	"fmt"
	"os"
)

type ConsoleLog struct {
	level       uint8
}

func NewConsoleLog(level string) Logger {
	log := new(ConsoleLog)
	log.level = getLogLevel(level)
	return log
}

func (f *ConsoleLog) Debug(format string, v ...interface{}) {
	thisLevel := getLogLevel("DEBUG")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *ConsoleLog) Error(format string, v ...interface{}) {
	thisLevel := getLogLevel("ERROR")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *ConsoleLog) Info(format string, v ...interface{}) {
	thisLevel := getLogLevel("INFO")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *ConsoleLog) Warn(format string, v ...interface{}) {
	thisLevel := getLogLevel("WARN")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *ConsoleLog) Fatal(format string, v ...interface{}) {
	thisLevel := getLogLevel("FATAL")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *ConsoleLog) setLog(level uint8, format string, v ...interface{})  {
	if f.level > level {
		return
	}
	logData := writeLog(level, format, v...)
	color := getLevelColor(level)
	fmt.Fprintf(os.Stdout, "%c[0;%dm %s %s %s %c[0m \n", ColorSeqClear, color, logData.TimeStr, logData.LevelStr, logData.Message, ColorSeqClear)
}
