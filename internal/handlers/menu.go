package handlers

import "github.com/gin-gonic/gin"

func Menu(c *gin.Context) {
	c.JSON(200, gin.H{"Message: ": "It´s ok"})
}