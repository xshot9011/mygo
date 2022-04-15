package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogConfiguration() {
	log.SetOutput(os.Stdout)

	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02T15:04:05+0000"
	customFormatter.FullTimestamp = true

	log.SetFormatter(customFormatter)
	log.SetLevel(logLevel)

	// The log level is the hierarchy.
	log.Trace("Successfully initial log Trace")
	log.Debug("Successfully initial log Debug")
	log.Info("Successfully initial log Info")
	log.Warn("Successfully initial log Warn")
	log.Error("Successfully initial log Error")
	// log.Fatal("Bye.")         // Calls os.Exit(1) after logging
	// log.Panic("I'm bailing.") // Calls panic() after logging
}
