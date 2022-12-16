package main

import (
	"api-rest-go-gin/controllers"
	"api-rest-go-gin/database"
	"api-rest-go-gin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	ID int
)

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Pedro Yoshimura",
		CPF:  "12864152924",
		RG:   "159357789",
	}

	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.GetConexao()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.Index)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestBuscaAlunoPorIdHandler(t *testing.T) {
	database.GetConexao()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.Show)

	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMock models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Pedro Yoshimura", alunoMock.Nome)
	assert.Equal(t, "12864152924", alunoMock.CPF)
	assert.Equal(t, "159357789", alunoMock.RG)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.GetConexao()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/document/:cpf", controllers.FindByDocument)

	req, _ := http.NewRequest("GET", "/alunos/document/12864152924", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.GetConexao()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.Update)

	aluno := models.Aluno{Nome: "Pedro Yoshimura", CPF: "12864152924", RG: "159357789"}
	valorJson, _ := json.Marshal(aluno)
	pathParaEditar := "/alunos/" + strconv.Itoa(ID)

	fmt.Println(pathParaEditar)

	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(response.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "Pedro Yoshimura", alunoMockAtualizado.Nome)
	assert.Equal(t, "12864152924", alunoMockAtualizado.CPF)
	assert.Equal(t, "159357789", alunoMockAtualizado.RG)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.GetConexao()
	CriaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.DELETE("alunos/:id", controllers.Delete)

	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDaBusca, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}
