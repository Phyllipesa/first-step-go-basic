package main

import (
	"fmt"
	"time"
)

/*
Nesse cenário a primeira fun tem seu funcionamento prejudicado pela segunda fun.

O canal 1 é lido primeiro e precisa de menos tempo, e o canal 2 é lido depois e precisa de mais tempo, então a primeira função fica limitada
ao tempo de execução da segunda.

Para resolver isso, é necessário usar o select, que permite que um canal
seja lido em paralelo a outro.

O select é uma estrutura de controle que permite aguardar múltiplos canais
em paralelo e executar a ação correspondente quando algum deles estiver pronto.
*/

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
