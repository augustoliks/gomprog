package main

import (
	"flag"
	"fmt"

	"github.com/augustoliks/pkg/gomprog/internal/incoming"
	"github.com/augustoliks/pkg/gomprog/internal/outgoing"
	"github.com/augustoliks/pkg/gomprog/internal/service"
)

func main() {
	fmt.Println("carlos")
	var (
		url      = flag.String("n", "127.0.0.1", "Graylog URL")
		user     = flag.String("u", "admin", "Graylog Username")
		password = flag.String("p", "admin", "Graylog Password")
	)
	rsyslog := incoming.Rsyslog{}
	redis := outgoing.Graylog{
		URL:      *url,
		User:     *user,
		Password: *password,
	}
	service.Run(rsyslog, redis)
}
