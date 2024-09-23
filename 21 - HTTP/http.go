package main

import (
	"log"
	"net/http"
)

// Função que sera chamado ao acessar a rota "/home"
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

// Função que sera chamado ao acessar a rota "/usuarios"
func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuários"))
}

func main() {
	// Rotas
	// URI - Identificador do Recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
