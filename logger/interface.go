//Powered By Jason Wang Author.
package logger

const (
	LOG_DEBUG int = iota
	LOG_INFO
	LOG_WARN
	LOG_ERROR
)

type Logger interface {
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Level() int
	SetLevel(l int)

	Prefix() string
	SetPrefix(prefix string)

	Flag() int
	SetFlag(flag int)
}
