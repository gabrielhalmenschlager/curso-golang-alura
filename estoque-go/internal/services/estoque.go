package services

import "github.com/gabrielhalmenschlager/curso-golang-alura/estoque-go/internal/models"

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
