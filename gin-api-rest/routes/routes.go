package routes

import (
	"github.com/gabrielhalmenschlager/curso-golang-alura/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.Run()
}
