package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/models"
	"github.com/gin-gonic/gin"
)

var pizzas []models.Pizza

func LoadPizzas() {
	file, err := os.Open("pizzaria/dados/pizzas.json")
	if err != nil {
		fmt.Println("Erro File:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error Decoding JSON:", err)
	}
}

func GetPizzas(c *gin.Context) {
	c.JSON(200, gin.H{
		"pizzas": pizzas,
	})
}

func GetPizzasByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": err.Error(),
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
