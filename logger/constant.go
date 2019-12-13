package logger

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
)

const (
	LogSplitTypeHour = iota
	LogSplitTypeSize
)

const ColorSeqClear = 0x1B
const (
	_            = iota + 30 // black
	FATAL
	INFO
	_
	_
	WARN
	ERROR
	DEBUG
)

func getLevelText(level uint8) string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	}
	return "UNKNOWN"
}
func getLevelColor(level uint8) int {
	switch level {
	case LogLevelDebug:
		return DEBUG
	case LogLevelInfo:
		return INFO
	case LogLevelWarn:
		return WARN
	case LogLevelError:
		return ERROR
	case LogLevelFatal:
		return FATAL
	}
	return INFO
}
func getLogLevel(level string) uint8 {
	switch level {
	case "DEBUG":
		return LogLevelDebug
	case "INFO":
		return LogLevelInfo
	case "WARN":
		return LogLevelWarn
	case "ERROR":
		return LogLevelError
	case "FATAL":
		return LogLevelFatal
	}
	return LogLevelDebug
}
