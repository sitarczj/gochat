package middleware

import (
	"gochat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authorized(c *gin.Context) {
	var authToken models.Token
	token := c.Request.Header["Token"]

	if token == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing auth token",
		})

		c.AbortWithStatus(http.StatusBadRequest)

		return
	}

	if err := models.DB.Where("Token = ?", token).First(&authToken).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid auth token",
		})

		c.AbortWithStatus(http.StatusBadRequest)

		return
	}
}
