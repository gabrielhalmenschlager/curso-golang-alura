package controllers

import (
	"strconv"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/models"
	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func GetPizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, p := range pizzas {
		if p.ID == id {
			c.JSON(200, p)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Pizza Not Found",
	})
}

func PostPizza(c *gin.Context) {
	var newPizza models.Pizza
	if err := c.ShouldBindJSON(&newPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	newPizza.ID = len(pizzas) + 1
	pizzas = append(pizzas, newPizza)
	SavePizza()
	c.JSON(201, newPizza)
}

func UpdatePizza(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	var updatedPizza models.Pizza
	if err := c.ShouldBindJSON(&updatedPizza); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	for i, p := range pizzas {
		if p.ID == id {
			pizzas[i] = updatedPizza
			pizzas[i].ID = id
			SavePizza()
			c.JSON(200, pizzas[i])
			return
		}
	}
	c.JSON(404, gin.H{"message": "Pizza Not Found"})
}

func DeletePizzaByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	for i, p := range pizzas {
		if p.ID == id {
			pizzas = append(pizzas[:i], pizzas[i+1:]...)
			SavePizza()
			c.JSON(200, gin.H{"message": "Pizza Deleted"})
		}
	}
	c.JSON(404, gin.H{"message": "Pizza Not Found"})
}
