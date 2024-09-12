package main

import "fmt"

func main() {
	fmt.Println("-----Maps-----")

	usuario := map[string]string{
		"nome":      "Lucas",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[int]string{
		1: "Lucas",
		2: "Silva",
	}
	fmt.Println(usuario2)

	usuario3 := map[string]int{
		"dia": 12,
		"mes": 9,
	}
	fmt.Println(usuario3)

	fmt.Println()
	fmt.Println("-----Aninhando Maps-----")

	usuario4 := map[string]map[string]string{
		"nomeCompleto": {
			"primeiro": "Lucas",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "UnB",
		},
	}

	fmt.Println(usuario4)

	fmt.Println()
	fmt.Println("-----Modificando Conteudo-----")

	delete(usuario4, "nomeCompleto")
	fmt.Println(usuario4)

	usuario4["signo"] = map[string]string{
		"nome": "Leao",
	}

	fmt.Println(usuario4)

}
