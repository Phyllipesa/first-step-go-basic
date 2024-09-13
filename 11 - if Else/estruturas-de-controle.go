package main

import "fmt"

func main() {

	fmt.Println("-----Estrutura de Controle-----")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	fmt.Println()
	fmt.Println("-----variavel if init-----")

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero")
	} else if outroNumero < -10 {
		fmt.Println("Numero é menor que zero")
	}

	fmt.Println()
}
