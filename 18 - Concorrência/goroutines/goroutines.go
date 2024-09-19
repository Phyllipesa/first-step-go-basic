package main

import (
	"fmt"
	"time"
)

// CONCORRÊNCIA != PARALELISMO
func main() {
	escrever("Start")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
