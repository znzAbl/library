package logger

import "testing"

func TestInitLogger(t *testing.T) {
	zlogger := InitLogger("FILE", "/Users/zhangnengzhi/sites", "DEBUG")
	zlogger.Debug("test file log %s", "test test test")
	zlogger.Info("test file log %s", "test test test")
	zlogger.Warn("test file log %s", "test test test")
	zlogger.Error("test file log %s", "test test test")
	zlogger.Fatal("test file log %s", "test test test")
}
