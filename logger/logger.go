package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func init() {
	log = logrus.New()
	log.SetOutput(os.Stdout)

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02T15:04:05+0000"
	customFormatter.FullTimestamp = true

	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	log.SetFormatter(customFormatter)
	log.SetLevel(logLevel)
}

func Trace(args ...interface{}) {
	log.Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	log.Tracef(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func GetCurrentLog() string {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	return fmt.Sprintf("%s.%d", f.Name(), line)
}
