package main

import (
	"fmt"

	"github.com/augustoliks/gomprog/internal/outgoing"
)

func main() {
	redis := outgoing.Graylog{"host", 1234, "usuario", "senha"}
	fmt.Println(redis)
}
