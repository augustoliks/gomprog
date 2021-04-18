package config

import (
	log "github.com/sirupsen/logrus"
)

func SetupLog(logLevel string) {
	var logrusLogLevel log.Level
	var err error
	logrusLogLevel, err = log.ParseLevel(logLevel)
	if err != nil {
		logrusLogLevel = log.ErrorLevel
	}
	log.SetLevel(logrusLogLevel)
}
