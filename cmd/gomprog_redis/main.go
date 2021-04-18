package main

import (
	"flag"

	"github.com/augustoliks/gomprog/pkg/config"
	"github.com/augustoliks/gomprog/pkg/incoming"
	"github.com/augustoliks/gomprog/pkg/osutils"
	"github.com/augustoliks/gomprog/pkg/outgoing"
	"github.com/augustoliks/gomprog/pkg/service"
)

func main() {
	var (
		url      = flag.String("n", "127.0.0.1:6379", "Redis URL")
		password = flag.String("p", "", "Redis Password")
		logLevel = flag.String("v", "fatal", "Log Level | trace, debug, info, warn, error, fatal, panic")
	)
	config.SetupLog(*logLevel)
	flag.Parse()

	go osutils.HandleSignals()

	rsyslog := incoming.Rsyslog{}
	redis := outgoing.RedisPlugin{
		URL:      *url,
		Password: *password,
	}
	service.MainLoop_PipelineLogs(rsyslog, redis)
}
