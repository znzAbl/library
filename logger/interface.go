//Powered By Jason Wang Author.
package logger

const (
	LOG_DEBUG int = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
)

type Logger interface {
	Debug(format string, v ...interface{})
	Error(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Fatal(format string, v ...interface{})
}
