package main

import (
	"fmt"
	"time"
)

// CONCORRÊNCIA != PARALELISMO
// A segunda chamada nunca acontecerá porque a primeira nunca terminará.
// O metodo "go" indica para o programa que, independente se a função terminopu ou não siga o fluxo do programa.
func main() {
	go escrever("Start") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
