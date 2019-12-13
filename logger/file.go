package logger

import (
	"fmt"
	"os"
	"time"
)

type FileLog struct {
	file        *os.File
	filePath    string
	date        string
	level       uint8
	logDataChan chan *LogData
}

func NewFileLog(filePath string, level string) Logger {
	log := new(FileLog)
	_, err := os.Stat(filePath)
	if nil != err {
		panic(fmt.Sprintf("logger file path err %s", err))
	}
	log.filePath = filePath
	log.level = getLogLevel(level)
	log.OpenFile()
	log.logDataChan = make(chan *LogData, 10000)
	go log.writeLogBackground()
	return log
}

func (f *FileLog) OpenFile() {

	date := time.Now().Format("2006-01-02")
	if f.date == date {
		return
	}
	filename := fmt.Sprintf("%s/%s.log", f.filePath, date)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if nil != err {
		if f.date == "" {
			panic(fmt.Sprintf("logger open file err %s", err))
		} else {
			f.Fatal("logger open file err %s", err)
		}
	}
	if f.date != "" {
		f.file.Close()
	}
	f.date = date
	f.file = file
}
func (f *FileLog) Debug(format string, v ...interface{}) {
	thisLevel := getLogLevel("DEBUG")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *FileLog) Error(format string, v ...interface{}) {
	thisLevel := getLogLevel("ERROR")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *FileLog) Info(format string, v ...interface{}) {
	thisLevel := getLogLevel("INFO")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *FileLog) Warn(format string, v ...interface{}) {
	thisLevel := getLogLevel("WARN")
	f.setLog(thisLevel, format, v...)
	return
}
func (f *FileLog) Fatal(format string, v ...interface{}) {
	thisLevel := getLogLevel("FATAL")
	f.setLog(thisLevel, format, v...)
	return
}

func (f *FileLog) setLog(Level uint8, format string, v ...interface{})  {
	if f.level > Level {
		return
	}
	logData := writeLog(Level, format, v...)
	select {
	case f.logDataChan <- logData:
	default:
	}
	return
}

func (f *FileLog) writeLogBackground()  {
	for logData := range f.logDataChan {
		f.OpenFile()
		color := getLevelColor(logData.Level)
		fmt.Fprintf(f.file, "%c[0;%dm %s %s %s %c[0m \n", ColorSeqClear, color, logData.TimeStr, logData.LevelStr, logData.Message, ColorSeqClear)
	}
}
