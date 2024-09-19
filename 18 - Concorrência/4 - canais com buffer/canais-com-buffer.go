package main

import (
	"fmt"
)

func main() {
	canal := make(chan string)
	canal <- "Olá mundo!" // deadlock!

	msg := <-canal
	fmt.Println(msg)
}
