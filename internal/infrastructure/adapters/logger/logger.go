package logger

import (
	"application/internal/domain/ports/logger"

	"github.com/sirupsen/logrus"
)

// Info(msg string, meta ...interface{})
// 	Error(msg string, meta error)
// 	Warn(msg string, meta ...interface{})
// 	Debug(msg string, meta ...interface{})

type logrusLog struct{}

func NewLogger() logger.Logger {
	return &logrusLog{}
}

func (*logrusLog) Error(msg string, meta error) {
	logrus.Error(msg, meta)
}

func (*logrusLog) Info(msg string, meta ...interface{}) {
	logrus.Info(msg, meta)
}

func (*logrusLog) Warn(msg string, meta ...interface{}) {
	logrus.Warn(msg, meta)
}

func (*logrusLog) Debug(msg string, meta ...interface{}) {
	logrus.Debug(msg, meta)
}
