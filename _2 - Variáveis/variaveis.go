package main

import "fmt"

func main() {

	//Declaração de variáveis
	var variavel1 string = "Variavel 1"

	//Declaração de variáveis com inferência de tipo
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	//Declaração de variáveis em bloco
	var (
		variavel3 string = "Hank"
		variavel4 string = "Midas"
	)

	fmt.Println(variavel3, variavel4)

	//Declaração de constantes
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//Declaração de variáveis em bloco
	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	//Troca de valores entre variáveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
