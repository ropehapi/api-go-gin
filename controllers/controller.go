package controllers

import (
	"api-rest-go-gin/database"
	"api-rest-go-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExibeAlunos(c *gin.Context){
	c.JSON(200, models.Alunos)
}

func CriaNovoAluno(c *gin.Context){
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()}) 
		return 
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}