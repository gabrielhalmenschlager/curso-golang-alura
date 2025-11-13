package controllers

import (
	"html/template"
	"net/http"

	"github.com/gabrielhalmenschlager/curso-golang-alura/alura_loja/models"
)

var temp = template.Must(template.ParseGlob("alura_loja/templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
