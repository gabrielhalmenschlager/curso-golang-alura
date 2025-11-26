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
		fmt.Println("Error File:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pizzas); err != nil {
		fmt.Println("Error Decoding JSON:", err)
	}
}

func SavePizza() {
	file, err := os.Create("pizzaria/dados/pizzas.json")
	if err != nil {
		fmt.Println("Error File:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(pizzas); err != nil {
		fmt.Println("Error Enconding JSON:", err)
	}
}

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
	c.JSON(200, gin.H{"method": "put"})
}

func DeletePizzaByID(c *gin.Context) {
	c.JSON(200, gin.H{"method": "delete"})
}
