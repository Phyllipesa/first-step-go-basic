package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá mundo"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// 1 canal de saida recebe informações vinda de 2 canais de entrada
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

// Escrever encapsula uma chamada para uma goroutine e retornar um canal
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
