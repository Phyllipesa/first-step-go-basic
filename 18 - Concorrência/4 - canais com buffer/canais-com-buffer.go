package main

import (
	"fmt"
)

// Canal (Channel) - É a forma de comunicação entre goroutines.
// Canal é um tipo de dado que permite a comunicação entre goroutines.
// Para criar um canal, usamos a palavra-chave `make(chan tipo)`.
// O canal é bufferizado, ou seja, pode armazenar valores em um buffer e enviar ou receber valores de um buffer.

// Exemplo de canal bufferizado
func main() {
	canal := make(chan string, 2)
	canal <- "Olá mundo!"
	canal <- "Programando em Go!"

	// Caso eu tente mandar ou receber outro valor acarretará em um deadlock, porque excediderá a quantidade do buffer

	msg := <-canal
	msg2 := <-canal

	fmt.Println(msg)
	fmt.Println(msg2)
}
