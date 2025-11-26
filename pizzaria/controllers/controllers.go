package controllers

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/models"
	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	pizzas := []models.Pizza{
		{ID: 1, Nome: "Toscana", Preco: 79.5},
		{ID: 2, Nome: "Marguerita", Preco: 69.5},
		{ID: 3, Nome: "Atum com queijo", Preco: 59.5},
	}
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}
