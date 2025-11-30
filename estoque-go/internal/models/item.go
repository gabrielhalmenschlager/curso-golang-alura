package models

import "fmt"

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

func (i Item) Info() string {
	return fmt.Sprintf("ID %d | Name: %s | Quantidade: %d | Price: %.2f", i.ID, i.Name, i.Quantity, i.Price)
}
