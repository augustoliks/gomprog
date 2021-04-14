package service

import "io"

type Incoming interface {
	ConfirmReceivedLogToRsyslog()
	ReceivedLog() (string, error)
	HandleSignals()
}

type Outgoing interface {
	OnInit()
	OnProcessed(string) string
	OnSend(string)
}

func Run(rsyslog Incoming, plugin Outgoing) {
	go rsyslog.HandleSignals()
	plugin.OnInit()

	for {
		logLine, err := rsyslog.ReceivedLog()
		if logLine == "" || err == io.EOF {
			break
		}
		logProcessed := plugin.OnProcessed(logLine)
		plugin.OnSend(logProcessed)
		rsyslog.ConfirmReceivedLogToRsyslog()
	}

}
