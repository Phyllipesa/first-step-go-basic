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

/*
	worker é uma função que recebe dois canais

tarefas é um canal ue só recebe valores
resultados é um canal que só envia valores

worker funciona como um pool de trabalhadores que recebem tarefas e enviam resultados
*/
func worker(tarefas <-chan int, resultado chan<- int) {
	for numero := range tarefas {
		resultado <- fibonacci(numero)
	}
}

/*
	 	fibonacci é uma função que recebe um número e retorna o valor da sequência
		de fibonacci na posição do número recebido como parâmetro.
*/
func fibonacci(position int) int {
	if position <= 1 {
		return position
	}
	return fibonacci(position-1) + fibonacci(position-2)
}
