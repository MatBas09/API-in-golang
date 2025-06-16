package handlers

import (
	"net/http"

	"github.com/MatBas09/API-in-golang/config"
	"github.com/MatBas09/API-in-golang/internal/models"
	"github.com/gin-gonic/gin"
)

func Users(c *gin.Context) {
	var allUsers []models.User

	result := config.DB.Find(&allUsers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Message": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, allUsers)
}

func CreatUsers(c *gin.Context) {
	user := models.User{Name: "juju_todinho", Email: "senhaforte@gmail.com", Password: "12345"}

	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Usuario criado com sucesso!"})
}