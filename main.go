package main

import (
	"github.com/sirupsen/logrus"
	"github.com/xshot9011/mygo/logger"
)

var log *logrus.Logger

func init() {
	log = logger.NewLogConfiguration()
}

func main() {
	log.Trace(logger.Trace())
	caller()
}

func caller() {
	log.Trace(logger.Trace())
}
