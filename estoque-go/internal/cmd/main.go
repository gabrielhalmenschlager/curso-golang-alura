package main

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/estoque-go/internal/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/estoque-go/internal/services"
)

func main() {
	fmt.Println("===== Sistema de Estoque =====")
	estoque := services.NewEstoque()
	itens := []models.Item{
		{ID: 1, Name: "Item 1", Quantity: 5, Price: 19.99},
		{ID: 2, Name: "Item 2", Quantity: 10, Price: 15.99},
		{ID: 3, Name: "Item 3", Quantity: 15, Price: 10.99},
	}
	for _, item := range itens {
		err := estoque.AddItem(item, "Gabriel")
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	for _, item := range estoque.ListItens() {
		fmt.Printf("\n ID: %d | Item: %s | Quantidade: %d | Preço %.2f", item.ID, item.Name, item.Quantity, item.Price)
	}

	fmt.Println()

	fmt.Printf("\nValor total do estoque: R$ %.2f", estoque.CalculateTotalCost())

	fmt.Println()

	logs := estoque.ViewAuditLog()
	for _, log := range logs {
		fmt.Printf("\n[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s", log.Timestamp.Format("01/02 15:04:05"), log.Action, log.User, log.ItemID, log.Quantity, log.Reason)
	}
}
