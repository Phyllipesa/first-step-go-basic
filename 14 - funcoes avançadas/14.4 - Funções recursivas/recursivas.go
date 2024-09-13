package main

import "fmt"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-1) + fibonacci(position-2)
}

func main() {
	fmt.Println("Funções Recursivas")
	// 1 1 2 3 5 8 13

	position := uint(12)
	fmt.Println(fibonacci(uint(position)))

	for i := uint(1); i <= position; i++ {
		fmt.Println(fibonacci(i))
	}
}
