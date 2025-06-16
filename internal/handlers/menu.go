package handlers

import (
	"net/http"

	"github.com/MatBas09/API-in-golang/config"
	"github.com/MatBas09/API-in-golang/internal/models"
	"github.com/gin-gonic/gin"
)

func Menu(c *gin.Context) {
	var menuItens []models.MenuItem

	result := config.DB.Find(&menuItens)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
	}
	c.JSON(http.StatusOK, menuItens)
}