package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup - garante que todas as goroutines terminem antes de encerrar o programa
// sync - pacote de sincronização para garantir a segurança de concorrência

/*
O WaitGroup é usado para esperar que todas as goroutines terminem antes de encerrar o programa.
Ele é útil quando você precisa aguardar a conclusão de várias operações concorrentes antes de prosseguir.
*/

/*
O método Add adiciona uma nova goroutine ao WaitGroup. Cada goroutine chamando Done decrementa o contador do WaitGroup.
O método Wait bloqueia até que o contador do WaitGroup seja zero, indicando que todas as goroutines terminaram.
*/

/*
Neste exemplo, criamos um WaitGroup e adicionamos duas goroutines. Cada goroutine chama a função escrever e,

	após completar sua execução, chama Done para decrementar o contador do WaitGroup.

O método Wait é chamado após adicionar as duas goroutines. O programa só será encerrado quando todas as goroutines terminarem.
*/
func main() {
	var waitGroup sync.WaitGroup

	// Adiciona uma goroutine ao WaitGroup
	waitGroup.Add(2)

	// Goroutine 1
	go func() {
		escrever("Olá!")
		waitGroup.Done() // Decrementa o contador do WaitGroup
	}()

	// Goroutine 2
	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // Decrementa o contador do WaitGroup
	}()

	// Espera até que todas as goroutines terminem
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
