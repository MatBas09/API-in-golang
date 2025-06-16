package routes

import (
	"github.com/MatBas09/API-in-golang/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(e *gin.Engine) {
	api := e.Group("/api")
	{
		api.GET("/menu", handlers.Menu)
		api.GET("/users", handlers.Users)
		api.POST("/creat/user", handlers.CreatUsers)
	}
}
