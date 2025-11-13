package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("alura_loja/templates/*.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, Bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
