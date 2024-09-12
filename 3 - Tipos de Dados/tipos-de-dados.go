package main

import (
	"errors"
	"fmt"
)

func main() {

	// Tipagem INT/ FLOAT/ STRING/ BOOL é associada por padrão
	// 	pela arquitetura do seu PC
	// exemplo: 32 / 64 bits

	// NUMEROS INTEIROS
	var numero int64 = -1000000000
	fmt.Println(numero)

	// Não aceita numeros negativos
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// Alias
	// INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	// BYTE = UNIT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// NUMEROS REAIS
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	// Retorna o numero do caractere na tabela ASCII
	char := 'B'
	fmt.Println(char)

	// BOOLEAN
	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
