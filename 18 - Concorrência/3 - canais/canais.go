package main

import (
	"fmt"
	"time"
)

/*
Canais são o jeito de fazer goroutines se comunicarem. Um canal é um tipo.
*/
func main() {
	// Criando um canal de tipo string(só poderá receber e enviar dados do tipo string)
	canal := make(chan string)
	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a se executada")

	// Recebendo a mensagem do canal, utilizar o <- antes do canal
	mensagem1 := <-canal
	fmt.Println(mensagem1)

	// Para cada mensagem que for recebida no canal(enquanto ele estiver aberto), ele printa na tela.
	// for mensagem := range canal {
	// 	fmt.Println(mensagem)
	// }

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// Enviando a mensagem para o canal, utilizar o <- depois do canal
		canal <- texto
		time.Sleep(time.Second)
	}

	// Deadlock, pois o canal está esperando receber uma mensagem, mas a função já terminou. Para resolver isso, fechar o canal.
	// close(canal)
}
