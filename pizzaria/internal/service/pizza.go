package service

import (
	"errors"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/models"
)

func ValidadePizzaPrice(pizza *models.Pizza) error {
	if pizza.Preco < 0 {
		return errors.New("o preço da pizza não pode ser negativo")
	}
	return nil
}
