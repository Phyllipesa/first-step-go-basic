package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	somar := somar(10, 20)
	fmt.Println(somar)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	// Ignorando um retorno
	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)

	_, resultadoSubtracao2 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSubtracao2)

	// Ignorando um retorno
}
