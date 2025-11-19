package main

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/api_rest_gin_go-aula_5/database"
	"github.com/gabrielhalmenschlager/curso-golang-alura/api_rest_gin_go-aula_5/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
