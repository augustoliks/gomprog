package service

import (
	"encoding/json"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type Incoming interface {
	NotifySuccess()
	ReceivedLog(stdin io.Reader) (string, error)
}

type Outgoing interface {
	OnInit() error
	OnSend(GELFLogFormat) error
}

type GELFLogFormat struct {
	Host         string `json:"host"`
	ShortMessage string `json:"short_message"`
	Timestamp    int    `json:"timestamp"`
	Group        string `json:"_group"`
	AppName      string `json:"_app_name"`
}

func convertedLog(logRaw string) (GELFLogFormat, error) {
	var gelfLog GELFLogFormat
	err := json.Unmarshal([]byte(logRaw), &gelfLog)
	return gelfLog, err
}

func MainLoop_PipelineLogs(rsyslog Incoming, plugin Outgoing) {
	log.Debugln("operation was started")

	err := plugin.OnInit()
	if err != nil {
		log.Fatal(err)
	} else {
		rsyslog.NotifySuccess()
	}

	for {
		logLine, err := rsyslog.ReceivedLog(os.Stdin)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}

		gelfLog, err := convertedLog(logLine)
		if err != nil {
			log.Warn("log discarted, log content is not compatible with gelf format | ", logLine)
			continue
		}

		err = plugin.OnSend(gelfLog)
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Debugln("log delivery with success | ", logLine)
			rsyslog.NotifySuccess()
		}
	}
}
