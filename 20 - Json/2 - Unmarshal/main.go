package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"vira-lata","idade":3}`

	var c cachorro

	// Convertendo de JSON para Struct
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	cachorro2EmJSON := `{"nome":"Tody","raca":"Poodle"}`

	c2 := make(map[string]string)

	// Convertendo de JSON para Map
	// Tomar cuidado com a tipagem, voce pode ter uma tipagem em INT e receber STRING
	// Exemplo: map[string]int e receber um dado do tipo `{"nome":"Tody","raca":"Poodle"}`
	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
