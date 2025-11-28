package main

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/data"
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", handler.GetPizzas)
	router.GET("/pizzas/:id", handler.GetPizzaByID)
	router.POST("/pizzas", handler.PostPizza)
	router.PUT("/pizzas/:id", handler.UpdatePizza)
	router.DELETE("/pizzas/:id", handler.DeletePizzaByID)
	router.Run()
}
