package outgoing

import (
	"github.com/augustoliks/gomprog/internal/service"
)

type Graylog struct {
	URL      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (graylog Graylog) OnInit() {}

func (graylog Graylog) OnSend(log service.GELFLogFormat) {}
