package controllers

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/models"
	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{
	{ID: 1, Nome: "Toscana", Preco: 79.5},
	{ID: 2, Nome: "Marguerita", Preco: 69.5},
	{ID: 3, Nome: "Atum com queijo", Preco: 59.5},
}

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func PostPizzas(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
		})
		return
	}
	pizzas = append(pizzas, newPizza)
}
