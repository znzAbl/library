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
		for f := 30; f <= 37; f++ { // 前景色彩 = 30-37
			for d := range []int{0, 1, 4, 5, 7, 8} { // 显示方式 = 0,1,4,5,7,8
				fmt.Printf(" %c[%d;%dm%s(d=%d,f=%d)%c[0m ", 0x1B, d, f, "", d, f, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	logData := writeLog(level, format, v...)
	color := getLevelColor(level)
	information := fmt.Sprintf("%c[0;37m[%s:%s:%d]%c[0m", ColorSeqClear, logData.Filename, logData.FuncName, logData.LineNo, ColorSeqClear)
	fmt.Fprintf(os.Stdout, "%c[0;%dm %s %s %s %c[0m \t %s \n", ColorSeqClear, color, logData.TimeStr, logData.LevelStr, logData.Message, ColorSeqClear, information)
}
