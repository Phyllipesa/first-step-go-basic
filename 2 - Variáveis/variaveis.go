package main

import "fmt"

func main() {

	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Hank"
		variavel4 string = "Midas"
	)

	fmt.Println(variavel3, variavel4)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
