package incoming

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var in = bufio.NewReader(os.Stdin)

type Rsyslog struct{}

func (rsyslog Rsyslog) HandleSignals() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)

	for s := range signals {
		switch s {
		case syscall.SIGUSR1, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt:
			os.Exit(1)
		}
	}
}

func (rsyslog Rsyslog) ReceivedLog() (string, error) {
	log_line, err := in.ReadString('\n')
	log_line = strings.TrimSuffix(log_line, "\n")

	if err != nil {
		return "", err
	}

	return log_line, nil
	// return `{"host":"carlos","short_message":"neto","timestamp": 123456,"_group": "123","_app_name":"123"}`, nil
}

func (rsyslog Rsyslog) ConfirmReceivedLogToRsyslog() {
	fmt.Println("OK")
}