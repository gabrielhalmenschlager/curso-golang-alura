package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/estoque-go/internal/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/estoque-go/internal/services"
)

func main() {
	fmt.Println("Sistema de Estoque")
	estoque := services.NewEstoque()
	itens := []models.Item{
		{ID: 1, Name: "Item 1", Quantity: 5, Price: 19.99},
		{ID: 2, Name: "Item 2", Quantity: 10, Price: 15.99},
		{ID: 3, Name: "Item 3", Quantity: 15, Price: 10.99},
	}
	for _, item := range itens {
		err := estoque.AddItem(item)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	fmt.Println(estoque)
}
