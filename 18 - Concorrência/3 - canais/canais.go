package main

import (
	"fmt"
	"time"
)

// Canal (Channel) - É a forma de comunicação entre goroutines.
// Canal é um tipo de dado que permite a comunicação entre goroutines.
// Para criar um canal, usamos a palavra-chave `make(chan tipo)`.
// O canal é bidirecional, ou seja, pode enviar e receber valores.
// O canal é bloqueante, ou seja, se o canal estiver cheio e você tentar enviar um valor, o programa vai bloquear até que algum valor seja recebido.
// O canal é unidirecional, ou seja, pode enviar valores mas não pode receber.
// O canal é fechado, ou seja, quando todos os valores foram enviados, o canal é fechado e não pode receber mais valores.
// O canal é bufferizado, ou seja, pode armazenar valores em um buffer e enviar ou receber valores de um buffer.
// O canal é sincronizado, ou seja, só uma goroutine pode enviar ou receber valores de um canal por vez.
// O canal é de tipo genérico, ou seja, pode armazenar valores de qualquer tipo.
func main() {
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a se executada")

	// Para cada mensagem que for recebida no canal(enquanto ele estiver aberto), ele printa na tela.
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
