package routes

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/pizzaria/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	controllers.LoadPizzas()
	router := gin.Default()
	router.GET("/pizzas", controllers.GetPizzas)
	router.GET("/pizzas/:id", controllers.GetPizzaByID)
	router.POST("/pizzas", controllers.PostPizza)
	router.PUT("/pizzas/:id", controllers.UpdatePizza)
	router.DELETE("/pizzas/:id", controllers.DeletePizzaByID)
	router.Run()
}
