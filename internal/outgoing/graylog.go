package outgoing

import "fmt"

type Graylog struct {
	host     string
	port     int
	user     string
	password string
}

func (g Graylog) OnInit() {
	fmt.Println("ok")
}

func (g Graylog) OnProcessed() {
	fmt.Println("ok")
}

func (g Graylog) OnSend() {
	fmt.Println("ok")
}
