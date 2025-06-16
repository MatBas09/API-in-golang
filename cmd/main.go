package main

import (
	"github.com/MatBas09/API-in-golang/config"
	"github.com/MatBas09/API-in-golang/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if config.DB == nil {
		config.ConnectDatabase()
	}

	e := gin.Default()
	routes.SetupRoutes(e)

	e.Run(":8080")
}