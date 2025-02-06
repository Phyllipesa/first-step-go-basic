package main

import (
	"fmt"
	"sync"
)

// Mutex para sincronizar o acesso ao cache
type SafeCache struct {
	mu    sync.Mutex
	cache map[int]int
}

// Instância global do cache protegido por mutex
var safeCache = SafeCache{cache: make(map[int]int)}

func main() {
	// Canal para enviar as tarefas (números de Fibonacci a serem calculados)
	tarefas := make(chan int, 45)
	// Canal para receber os resultados dos cálculos
	resultados := make(chan int, 45)
	// WaitGroup para aguardar a finalização dos workers
	var wg sync.WaitGroup

	// Criando 4 workers concorrentes
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(tarefas, resultados, &wg)
	}

	// Enviando 45 tarefas para cálculo de Fibonacci
	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	// Fecha o canal de tarefas, indicando que não há mais trabalho
	close(tarefas)

	// Aguarda os workers terminarem antes de fechar o canal de resultados
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Lendo e imprimindo os resultados conforme chegam dos workers
	for resultado := range resultados {
		fmt.Println(resultado)
	}
}

// Worker: consome tarefas do canal e envia os resultados
func worker(tarefas <-chan int, resultados chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Indica que este worker terminou

	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

// fibonacci calcula o número de Fibonacci na posição fornecida
// Usa um cache seguro com mutex para evitar recomputação desnecessária
func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	// Verifica se o valor já foi calculado e está armazenado no cache
	safeCache.mu.Lock()
	val, exists := safeCache.cache[position]
	safeCache.mu.Unlock()

	if exists {
		return val
	}

	// Calcula recursivamente os valores de Fibonacci
	result := fibonacci(position-1) + fibonacci(position-2)

	// Armazena o resultado no cache para evitar cálculos repetidos
	safeCache.mu.Lock()
	safeCache.cache[position] = result
	safeCache.mu.Unlock()

	return result
}
