package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultado chan<- int) {
	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}
}

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonacci(position-1) + fibonacci(position-2)
}
