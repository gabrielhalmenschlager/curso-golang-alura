package routes

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	controllers.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", controllers.GetPizzas)
	router.GET("/pizzas/:id", controllers.GetPizzasByID)
	router.POST("/pizzas", controllers.PostPizzas)
	router.Run()
}
