package main

import "fmt"

func main() {
	// Funções anonimas
	// Função sem parâmetros
	func() {
		fmt.Println("Olá Mundo!")
	}()

	// Função com parâmetros
	func(text string) {
		fmt.Println(text)
	}("Hello World!")

	// Função com parâmetros e retorno
	resultado := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Hello World!")

	fmt.Println(resultado)

	retorno := func(text string) string {
		return fmt.Sprintf("Recebido -> %s %d", text, 10)
	}("Hello World!")

	fmt.Println(retorno)
}
