package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Pessoa struct {
	Nome  string `json:"name"`
	Idade int    `json:"age"`
}

func main() {
	pessoaRaw := []byte(`{"name":"carlos neto","age":23}`)

	var pessoaLoad Pessoa

	json.Unmarshal(pessoaRaw, &pessoaLoad)

	fmt.Println(pessoaLoad)

	pessoa := Pessoa{
		Nome:  "carlos",
		Idade: 23,
	}

	pessoa.Nome = "carlos neto"
	jsonData, err := json.Marshal(pessoa)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(jsonData))
}
