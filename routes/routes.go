package routes

import (
	"api-rest-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.Index)
	r.GET("/alunos/:id", controllers.Show)
	r.GET("/alunos/document/:cpf", controllers.FindByDocument)
	r.POST("/alunos", controllers.Store)
	r.PATCH("/alunos/:id", controllers.Update)
	r.DELETE("alunos/:id", controllers.Delete)
	r.Run()
}
