package service

import (
	"encoding/json"
	"io"
)

type Incoming interface {
	ConfirmReceivedLogToRsyslog()
	ReceivedLog() (string, error)
	HandleSignals()
}

type Outgoing interface {
	OnInit()
	OnSend(GELFLogFormat)
}

type GELFLogFormat struct {
	Host         string `json:"host"`
	ShortMessage string `json:"short_message"`
	Timestamp    string `json:"timestamp"`
	Group        string `json:"_group"`
	AppName      string `json:"_app_name"`
}

func convertedLog(logRaw string) GELFLogFormat {
	var gelfLog GELFLogFormat
	json.Unmarshal([]byte(logRaw), &gelfLog)
	return gelfLog
}

func Run(rsyslog Incoming, plugin Outgoing) {
	go rsyslog.HandleSignals()
	plugin.OnInit()

	for {
		logLine, err := rsyslog.ReceivedLog()
		if logLine == "" || err == io.EOF {
			break
		}
		gelfLog := convertedLog(logLine)
		plugin.OnSend(gelfLog)
		rsyslog.ConfirmReceivedLogToRsyslog()
	}
}
