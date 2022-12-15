package routes

import (
	"api-rest-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeAlunos)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.Run()
}
