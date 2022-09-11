package logger

type Logger interface {
	Info(msg string, meta ...interface{})
	Error(msg string, meta error)
	Warn(msg string, meta ...interface{})
	Debug(msg string, meta ...interface{})
}
