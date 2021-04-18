package config

import (
	"github.com/sirupsen/logrus"
)

func SetupLog(logLevel string) {
	var logrusLogLevel logrus.Level
	var err error
	logrusLogLevel, err = logrus.ParseLevel(logLevel)
	if err != nil {
		logrusLogLevel = logrus.ErrorLevel
	}
	logrus.SetLevel(logrusLogLevel)
}
