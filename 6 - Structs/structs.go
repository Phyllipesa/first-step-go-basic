package main

import "fmt"

// struct é um tipo de dado que permite armazenar diferentes tipos de dados
type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	var u usuario
	u.nome = "Bender"
	u.idade = 55
	fmt.Println(u)

	enderecoEx := endereco{"Avenida 51", 0}

	// Forma de declarar um struct com atributos específicos
	usuario4 := usuario{"Arthur", 30, enderecoEx}
	fmt.Println(usuario4)

	usuario2 := usuario{nome: "Arthur", endereco: enderecoEx}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 18}
	fmt.Println(usuario3)
}
