package main

import (
	"fmt"
	"time"
)

/*
	CONCORRÊNCIA != PARALELISMO

A segunda chamada nunca acontecerá porque a primeira nunca terminará.
O metodo "go" indica para o programa que, independente se a função
terminou ou não siga o fluxo do programa.
*/
func main() {
	// goroutine tem a função de executar uma função em paralelo com o programa principal.
	go escrever("Start") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
