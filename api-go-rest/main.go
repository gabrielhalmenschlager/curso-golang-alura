package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/api-go-rest/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/api-go-rest/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
