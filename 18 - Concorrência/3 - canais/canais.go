package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a se executada")

	mensagem := <-canal
	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
}
