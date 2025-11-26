package routes

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()
	router.GET("/pizzas", controllers.GetPizzas)
	router.POST("/pizzas", controllers.PostPizzas)
	router.Run()
}
