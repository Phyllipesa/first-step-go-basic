// TESTE DE UNIDADE

package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TestTipoDeEndereco(t *testing.T) {

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
