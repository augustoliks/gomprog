package outgoing

import (
	"fmt"
)

type Redis struct {
	host     string
	port     int
	user     string
	password string
}

func (redis Redis) OnInit() {
	fmt.Println("ok")
}

func (redis Redis) OnProcessed() {
	fmt.Println("ok")
}

func (redis Redis) OnSend() {
	fmt.Println("ok")
}
