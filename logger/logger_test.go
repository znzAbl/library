package logger

import (
	"testing"
)

//TestFileLogger：12
//Debug：79
//writeLog：68
//GetLineInfo:10
//runtime.Caller(3)
func TestFileLogger(t *testing.T) {
	LogLevelDebug := make(map[string]string)
	logger, _ := NewFileLogger(LogLevelDebug)
	logger.Debug("user id[%d] is come from china", 324234)
	logger.Warn("test warn log")
	logger.Fatal("test fatal log")
	logger.Close()
}

//TestFileLogger：12
//Debug：79
//writeLog：68
//GetLineInfo:10
//runtime.Caller(3)
func TestConsoleLogger(t *testing.T) {
	LogLevelDebug := make(map[string]string)
	logger, _ := NewConsoleLogger(LogLevelDebug)
	logger.Debug("user id[%d] is come from china", 324234)
	logger.Warn("test warn log")
	logger.Fatal("test fatal log")
	logger.Close()
}
