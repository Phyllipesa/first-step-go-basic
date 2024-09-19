package main

import (
	"fmt"
	"time"
)

/*
Nesse cenário a primeira fun tem seu funcionamento prejudicado pela segunda fun.

O canal 1 é lido primeiro e precisa de menos tempo, e o canal 2 é lido depois e precisa de mais tempo, então a primeira função fica limitada
ao tempo de execução da segunda.
*/
func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		mensagemCanal1 := <-canal1
		fmt.Println(mensagemCanal1)

		mensagemCanal2 := <-canal2
		fmt.Println(mensagemCanal2)
	}
}
