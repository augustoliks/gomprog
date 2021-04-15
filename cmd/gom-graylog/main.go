package main

import (
	"flag"

	"github.com/augustoliks/gomprog/internal/incoming"
	"github.com/augustoliks/gomprog/internal/outgoing"
	"github.com/augustoliks/gomprog/internal/service"
)

func ui(url *string, user *string, password *string) {
	flag.StringVar(url, "n", "127.0.0.1", "Graylog URL")
	flag.StringVar(user, "u", "admin", "Graylog Username")
	flag.StringVar(password, "p", "admin", "Graylog Password")
	flag.Parse()
}

func main() {
	var url string
	var user string
	var password string

	ui(&url, &user, &password)

	rsyslog := incoming.Rsyslog{}
	redis := outgoing.Graylog{
		URL:      url,
		User:     user,
		Password: password,
	}
	service.Run(rsyslog, redis)
}
