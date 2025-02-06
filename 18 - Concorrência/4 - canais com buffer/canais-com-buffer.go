package main

import (
	"fmt"
)

/*
	Canal com buffer

O canal é bufferizado, ou seja, pode armazenar valores em um buffer e enviar ou receber valores de um buffer.

Exemplo de canal bufferizado
*/
func main() {
	/*
		Nesse exemplo quando executamos o código, o programa irá travar e exibirá o erro fatal: all goroutines are asleep - deadlock!.
		Isso ocorre porque o canal está bloqueado, esperando que um valor seja enviado para ele, mas não há nenhuma goroutine para
		enviar um valor para o canal. Para corrigir esse erro, precisamos enviar um valor para o canal em uma goroutine.

		No caso, o programa ficará travado na linha canal <- "Olá mundo!" esperando algum valor, logo ele nunca chegará a linha
		mensagem := <-canal, pois o programa está travado na linha anterior.

		canal := make(chan string)
		canal <- "Olá mundo!"

		mensagem := <-canal
		fmt.Println(mensagem)
	*/

	/*
		Para corrigir o erro, podemos utilizar o buffer do canal, que é um espaço de armazenamento temporário para valores enviados
		para o canal. O buffer é definido ao criar o canal e é o segundo argumento da função make. No exemplo abaixo, o canal tem um
		buffer de 2, o que significa que pode armazenar até 2 valores. Quando o canal está cheio, ele bloqueia a goroutine que tenta
		enviar um valor para o canal. Quando o canal está vazio, ele bloqueia a goroutine que tenta receber um valor do canal.
	*/
	canal := make(chan string, 2)
	canal <- "Olá mundo!"
	canal <- "Programando em Go!"

	// Caso eu tente mandar ou receber outro valor acarretará em um deadlock, porque excediderá a quantidade do buffer

	msg := <-canal
	msg2 := <-canal

	fmt.Println(msg)
	fmt.Println(msg2)
}
