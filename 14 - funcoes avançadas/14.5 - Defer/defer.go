package main

import "fmt"

func fun1() {
	fmt.Println("Executando a função 1")
}

func fun2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer fmt.Println("Resultado do calculo:")
	fmt.Println("Calculando se o aluno está aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	defer fun1()
	fun2()
	fmt.Println(alunoEstaAprovado(7, 8))
}
