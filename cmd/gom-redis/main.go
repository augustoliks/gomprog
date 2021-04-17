package main

import (
	"flag"

	"github.com/augustoliks/gomprog/internal/incoming"

	"github.com/augustoliks/gomprog/internal/outgoing"
	"github.com/augustoliks/gomprog/internal/service"
)

func main() {
	var (
		url      = flag.String("n", "127.0.0.1:6379", "Redis URL")
		password = flag.String("p", "", "Redis Password")
	)

	flag.Parse()

	rsyslog := incoming.Rsyslog{}
	redis := outgoing.RedisPlugin{
		URL:      *url,
		Password: *password,
	}
	service.Run(rsyslog, redis)
}
