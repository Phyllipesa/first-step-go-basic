package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint64 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	// Converte o corpo da requisição para um struct do tipo usuario
	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o corpo da requisição em struct!"))
		return
	}

	// Conecta no banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	// Preparando o comando de inserção
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	// Inserindo as informações no DB
	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	// Recuperando o ID do usuario inserido no banco
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido!"))
		return
	}

	// Retornando o ID do usuário inserido com sucesso
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

	fmt.Println(usuario)
}

// BuscarUsuarios busca todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Conecta no banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	// Retornandp as linhas
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer db.Close()

	// Criando uma array de usuarios
	// Iterando por cada uma das linhas e criando um usuario
	// Adicionando cada usuario ao array usuarios
	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear os usuários!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	// Retornando os usuários em formato JSON
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuários em JSON!"))
		return
	}
}

// BuscarUsuario busca um usuário específico no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
