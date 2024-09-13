package main

import "fmt"

// Função recebe qualquer quantidade de parametros.
func soma(numeros ...int) {
	fmt.Println(numeros)
}

func soma2(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	// Função com parâmetros variáveis.
	soma(1, 2, 3)
	soma(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println("----- -----")

	// Função com retorno.
	totalSoma := soma2(1, 2, 3, 4, 5, 6, 200, 102, 12, 13)
	fmt.Println(totalSoma)
	fmt.Println("----- -----")

	// Função com vários parâmetros.
	escrever("Olá Mundo", 1, 2, 3, 4)
}
