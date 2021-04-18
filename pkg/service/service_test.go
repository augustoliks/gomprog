package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertedLog_LogRawWithGELFValidFormat(t *testing.T) {
	logExpected := GELFLogFormat{
		Host:         "localhost",
		ShortMessage: "fake log",
		Timestamp:    1618689190,
		Group:        "cluster-a",
		AppName:      "gomprog",
	}
	logRaw := `{"host":"localhost","short_message":"fake log","timestamp": 1618689190,"_group": "cluster-a","_app_name":"gomprog"}`
	logGelf, err := convertedLog(logRaw)

	assert.NoError(t, err)
	assert.Equal(t, logGelf, logExpected)
}

func TestConvertedLog_LogRawWithGELFInvalidFormat(t *testing.T) {
	var err error

	logInvalids := []string{
		`{"host": "bug""localhost","short_message":"fake log","timestamp": 1618689190,"_group": "cluster-b","_app_name":"gomprog"}`,
		`"{"host": "localhost","short_message":"fake log","timestamp": 1618689190,"_group": "cluster-b","_app_name":"gomprog"}"`,
	}

	for _, logInvalid := range logInvalids {
		_, err = convertedLog(logInvalid)
		assert.Error(t, err)
	}
}
