package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo")

	for i := 0; i < 3; i++ {
		fmt.Println(<-canal)
	}
}

// Escrever encapsula uma chamada para uma goroutine e retornar um canal
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
