package main

import "fmt"

func main() {
	fmt.Println("-----ARRAYS-----")
	// Array é um tipo de dado que armazena uma coleção de valores do mesmo tipo.
	// O tamanho do array é fixo e está definido na criação do array.
	// O array é uma estrutura de dados homogênea (todos os elementos têm o mesmo tipo).
	// O array é imutável, ou seja, não pode ser modificado após ser criado.

	// Declaração de um array com tamanho fixo
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// Declaração de um array com tamanho variável
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	fmt.Println("-----SLICES-----")

	// SLICE
	slice := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice)

	// Adicionando elementos ao slice
	slice = append(slice, 17)
	fmt.Println(slice)

	// Alterando o valor da posição 0 do slice
	slice[0] = 0
	fmt.Println(slice)

	// Criando um slice a partir de um array
	// OBS: O valor do indice 3 não esta incluido, ele só inclui o valor do indice anterior ao especificado
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	// O array original é alterado, pois o slice faz uma referência ao array original
	fmt.Println(slice2)

	// ARRAYS INTERNOS
	fmt.Println("-----ARRAYS INTERNOS-----")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 5)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity

	// O Slice estava com a sua capacidade esgotada(11/11). Quando o GO entende que o array vai "estourar"
	//  ele cria um novo array para se referenciar. Se o array tinha 11 elementos e eu estou querendo
	// colocar 12, então ele cria uma array de 24.
	slice3 = append(slice3, 6)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity
}
