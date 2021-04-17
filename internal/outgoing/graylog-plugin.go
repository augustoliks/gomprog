package outgoing

import (
	"github.com/augustoliks/gomprog/internal/service"
)

type Graylog struct {
	URL      string
	User     string
	Password string
}

func (graylog Graylog) OnInit() {}

func (graylog Graylog) OnSend(log service.GELFLogFormat) {}
