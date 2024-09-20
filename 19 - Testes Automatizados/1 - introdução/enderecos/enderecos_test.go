// TESTE DE UNIDADE

package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua Amadeus"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRetornado := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRetornado != tipoDeEnderecoEsperado {
		t.Errorf(
			"O tipo de endereço retornado é diferente do esperado! Esperava %s e recebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRetornado,
		)
	}
}
