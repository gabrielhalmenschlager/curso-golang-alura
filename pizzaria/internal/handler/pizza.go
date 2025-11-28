package handler

import (
	"strconv"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/data"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/models"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/service"
	"github.com/gin-gonic/gin"
)

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": data.Pizzas,
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
	for _, p := range data.Pizzas {
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
	if err := service.ValidadePizzaPrice(&newPizza); err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	newPizza.ID = len(data.Pizzas) + 1
	data.Pizzas = append(data.Pizzas, newPizza)
	data.SavePizza()
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
	if err := service.ValidadePizzaPrice(&updatedPizza); err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas[i] = updatedPizza
			data.Pizzas[i].ID = id
			data.SavePizza()
			c.JSON(200, data.Pizzas[i])
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
	for i, p := range data.Pizzas {
		if p.ID == id {
			data.Pizzas = append(data.Pizzas[:i], data.Pizzas[i+1:]...)
			data.SavePizza()
			c.JSON(200, gin.H{"message": "Pizza Deleted"})
		}
	}
	c.JSON(404, gin.H{"message": "Pizza Not Found"})
}
