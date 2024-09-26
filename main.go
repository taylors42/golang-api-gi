package main

import (
	"api-gin/database"
	"api-gin/routes"
	// "github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDb()
	routes.HandleRequests()
}
