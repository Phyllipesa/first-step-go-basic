package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Jack", 26}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "Cone Island"}
	fmt.Println(e1)
}
