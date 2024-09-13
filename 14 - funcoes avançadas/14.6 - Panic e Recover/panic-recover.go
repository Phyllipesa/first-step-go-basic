package main

import "fmt"

// panic: interrompe o fluxo de execução do programa
// recover: captura o panic e recupera o fluxo de execução do programa

// defer: adiar a execução de uma função até o final do escopo atual

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

// SEM RECOVER
func alunoEstaAprovado(n1, n2 float64) bool {
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")
}

// COM RECOVER
func alunoEstaAprovado2(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	// fmt.Println(alunoEstaAprovado2(8, 6))
	fmt.Println("Pós execução!")
}
