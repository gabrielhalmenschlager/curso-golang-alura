package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/estoque-go/internal/services"
)

func main() {
	fmt.Println("Sistema de Estoque")
	estoque := services.NewEstoque()
	fmt.Println(estoque)
}
