package main

import "fmt"

// Passando parametro por copia
func inverterSinal(number int) int {
	return number * -1
}

// Passando parametro por referência
func inverterSinalPonteiro(number *int) {
	*number = *number * -1
}

func main() {
	// Alterando o valor da copia
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	// Alterando o valor da variável original
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
