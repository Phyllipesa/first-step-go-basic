package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

// Função que sera chamado ao acessar a rota "/home"
func home(w http.ResponseWriter, r *http.Request) {
	usuario := usuario{"João", "email@e.com.br"}
	templates = template.Must(template.ParseGlob("*.html"))
	templates.ExecuteTemplate(w, "home.html", usuario)
}

func main() {

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
