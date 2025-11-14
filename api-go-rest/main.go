package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/api-go-rest/routes"
)

func main() {
	fmt.Println("iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
