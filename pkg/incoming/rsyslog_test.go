package incoming

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReceivedLog(t *testing.T) {
	rsyslog := Rsyslog{}
	var stdin bytes.Buffer

	logReceivedFromRsyslog := `{"host":"localhost","short_message":"fake log","timestamp": 1618689190,"_group": "cluster-a","_app_name":"gomprog"}`
	stdin.Write([]byte(logReceivedFromRsyslog))

	result, err := rsyslog.ReceivedLog(&stdin)

	assert.Error(t, err, io.EOF)
	assert.Equal(t, result, logReceivedFromRsyslog)
}
