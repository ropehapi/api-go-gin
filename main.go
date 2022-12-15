package main

import (
	"api-rest-go-gin/database"
	"api-rest-go-gin/routes"
)

func main() {
	database.GetConexao()
	routes.HandleRequests()
}
