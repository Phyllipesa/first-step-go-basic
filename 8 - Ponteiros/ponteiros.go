package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REREFÊNCIA DE MEMÓRIA
	// O *int é a tipagem para armazenamento de endereço de memoria
	var variavel3 int
	var ponteiro *int
	fmt.Println(variavel3, ponteiro)

	// Para acessar e copiar o endereço de memória da variavel3, precisamos usar o &
	variavel3 = 100
	ponteiro = &variavel3

	// Mostra o valor da variavel3
	// Mostra o endereço de memória da variavel3
	fmt.Println(variavel3, ponteiro)

	// Acessando o valor da variavel3 através do ponteiro
	fmt.Println(variavel3, *ponteiro)

}
