package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem o ipo válido e retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	// RUA DOS BOBOS COMPLEMENTO ALGUMA COISA
	// ["RUA", "DOS", "BOBOS", "COMPLEMENTO", "ALGUMA", "COISA"]
	enderecoEmLetrasMinusculas := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetrasMinusculas, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo inválido"
}
