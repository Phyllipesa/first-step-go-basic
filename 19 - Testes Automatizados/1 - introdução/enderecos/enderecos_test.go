// TESTE DE UNIDADE

package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// go test ./... = esse comando faz com que o go execute todos os testes do projeto
// go test -v = esse comando mostra mais informações durante o teste
// go test --cover = exibe a cobertura dos seus testes
// go test --coverprofile nome-do-arquivo.txt = Faz um arquivo contendo um relatório sobre as linhas que estão cobertas ou não pelos testes
// go tool cover --func=cobertura.txt = Mostra todas as funções que estão sendo testadas de acordo com o relatório.
// go tool cover --html=cobertura.txt = Mostra em um arquivo HTML todas as linhas que não estão cobertas por testes.

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TestTipoDeEndereco(t *testing.T) {
	// t.Parallel() = faz com que o teste seja executado paralelamente com outros
	// Fica na primeira linha e precisa ser colocado em todo o teste que executará paralelamente

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Praça das Rosas", "Tipo inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo inválido"},
	}

	//range -> percorre o array
	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf(
				"O tipo recebido é diferente do esperado! Esperava %s e recebi %s",
				cenario.retornoEsperado,
				retornoRecebido,
			)
		}
	}
}
