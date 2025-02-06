package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	/*
		os.Args é uma variável que contém todos os argumentos passados para o programa
		Se o programa for executado sem argumentos, os.Args terá apenas um elemento, o nome do programa
		Se o programa for executado com argumentos, os.Args terá todos os argumentos passados
		Por exemplo, se o programa for executado com "go run main.go ip --host devbook.com.br",
		 os.Args será igual a []string{"main.go", "ip", "--host", "devbook.com.br"}
		Usado para reconhecer os argumentos passados para o programa via linha de comando
	*/
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
