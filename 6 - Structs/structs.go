package main

import "fmt"

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

	usuario2 := usuario{nome: "Arthur", endereco: enderecoEx}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 18}
	fmt.Println(usuario3)
}
