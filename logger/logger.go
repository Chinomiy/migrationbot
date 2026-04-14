package logger

import (
	"os"

	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func Init() {
	logger, _ := zap.NewProduction()
	log = logger.Sugar()
}

func Info(msg string) {
	log.Info(msg)
}

func Infof(msg string, args ...interface{}) {
	log.Infof(msg, args...)
}

func Error(err error) {
	if err == nil {
		return
	}
	log.Error(err)
}

func Fatal(err error) {
	log.Error(err)
	os.Exit(1)
}

func Warn(msg string) {
	log.Warn(msg)
}

func Warnf(format string, args ...any) {
	log.Warnf(format, args...)
}
