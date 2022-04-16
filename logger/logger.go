package logger

import (
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

	// The log level is the hierarchy.
	// log.Trace("Successfully initial log Trace")
	// log.Debug("Successfully initial log Debug")
	// log.Info("Successfully initial log Info")
	// log.Warn("Successfully initial log Warn")
	// log.Error("Successfully initial log Error")
	// log.Fatal("Bye.")         // Calls os.Exit(1) after logging
	// log.Panic("I'm bailing.") // Calls panic() after logging
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
	log.Fatal(args...)
}
func Panicf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func LogFunction() {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	_, line := f.FileLine(pc[0])
	log.Tracef("Invoke at %s.%d", f.Name(), line)
}
