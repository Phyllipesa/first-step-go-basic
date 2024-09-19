package main

import (
	"fmt"
	"time"
)

// CONCORRÊNCIA != PARALELISMO
// A segunda chamada nunca acontecerá porque a primeira nunca terminará.
func main() {
	escrever("Start")
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
