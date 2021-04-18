package incoming

import (
	"bufio"
	"fmt"
	"io"
)

type Rsyslog struct{}

var in *bufio.Reader

func (rsyslog Rsyslog) ReceivedLog(stdin io.Reader) (string, error) {
	if in == nil {
		in = bufio.NewReader(stdin)
	}
	logLine, err := in.ReadString('\n')
	return logLine, err
}

func (rsyslog Rsyslog) NotifySuccess() {
	fmt.Println("OK")
}
