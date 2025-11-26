package main

import (
	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	Nome  string
	Preco float64
}

func main() {
	router := gin.Default()
	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": "Calabrasa",
		})
	})

	// nomePizzaria := "Pizzaria Go"
	// instagram, telefone := "@pizzaria_go", 555199999999

	// pizzas := []Pizza{
	// 	{ID: 1, Nome: "Toscana", Preco: 79.5},
	// 	{ID: 2, Nome: "Marguerita", Preco: 69.5},
	// 	{ID: 3, Nome: "Atum com queijo", Preco: 59.5},
	// }

	// fmt.Println(nomePizzaria, instagram, telefone)
	// fmt.Println(pizzas)

	router.Run()
}
