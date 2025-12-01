package services

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gabrielhalmenschlager/curso-golang-alura/estoque-go/internal/models"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (e *Estoque) AddItem(item models.Item, user string) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar item: [ID:%d] possui uma quantidade inválida (zero ou negativa)", item.ID)
	}
	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}
	e.items[strconv.Itoa(item.ID)] = item

	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Entrada de estoque",
		User:      user,
		ItemID:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando novos itens no estoque",
	})

	return nil
}

func (e *Estoque) ListItens() []models.Item {
	var itemList []models.Item
	for _, item := range e.items {
		itemList = append(itemList, item)
	}
	return itemList
}

func (e *Estoque) ViewAuditLog() []models.Log {
	return e.logs
}

func (e *Estoque) CalculateTotalCost() float64 {
	var totalCost float64
	for _, item := range e.items {
		totalCost += float64(item.Quantity) * item.Price
	}
	return totalCost
}

func (e *Estoque) RemoveItem(itemID int, quantity int, user string) error {
	existingItem, exists := e.items[strconv.Itoa(itemID)]
	if !exists {
		return fmt.Errorf("erro ao remover item: [ID:%d] não existe no estoque", itemID)
	}

	if quantity <= 0 {
		return fmt.Errorf("erro ao remover item: quantidade inválida (zero ou negativa) para [ID:%d]", itemID)
	}

	if existingItem.Quantity < quantity {
		return fmt.Errorf("erro ao remover item: estoque insuficiente para [ID:%d]. Disponível: %d, Solicitado: %d", itemID, existingItem.Quantity, quantity)
	}

	existingItem.Quantity -= quantity
	if existingItem.Quantity == 0 {
		delete(e.items, strconv.Itoa(itemID))
	} else {
		e.items[strconv.Itoa(itemID)] = existingItem
	}

	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Saída de estoque",
		User:      user,
		ItemID:    itemID,
		Quantity:  quantity,
		Reason:    "Removendo itens do estoque",
	})

	return nil
}
