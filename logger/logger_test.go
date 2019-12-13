package logger

import "testing"

func TestInitLogger(t *testing.T) {
	InitLogger("FILE", "/Users/zhangnengzhi/sites", "DEBUG")
	Log.Debug("test file log %s", "test test test")
	Log.Info("test file log %s", "test test test")
	Log.Warn("test file log %s", "test test test")
	Log.Error("test file log %s", "test test test")
	Log.Fatal("test file log %s", "test test test")
}
