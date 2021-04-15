package outgoing

import (
	"github.com/augustoliks/gomprog/internal/service"
)

type Redis struct {
	URL      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (redis Redis) OnInit() {}

func (redis Redis) OnSend(log service.GELFLogFormat) {}
