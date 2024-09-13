package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Loops-----")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println()
	// fmt.Println("-----Loops 2-----")

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println()
	// fmt.Println("-----Loops 3-----")

	// nomes := [3]string{"Dagni", "Hearden", "Francisco"}

	// for indice, nome := range nomes {
	// 	fmt.Println(indice, nome)
	// }

	// for nome := range nomes {
	// 	fmt.Println(nome)
	// }

	// for _, nome := range nomes {
	// 	fmt.Println(nome)
	// }

	// fmt.Println()
	// fmt.Println("-----Loops 4-----")

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, letra)
	// }

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// }

	fmt.Println()
	fmt.Println("-----Loops 5-----")

	usuario := map[string]string{
		"nome":      "Dagni",
		"sobrenome": "Taggart",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	fmt.Println()
	fmt.Println("-----Loops 6-----")

	// for {
	// 	fmt.Println("Executando infinitamente")
	// 	time.Sleep(time.Second)
	// }
}
