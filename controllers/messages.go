package controllers

import (
	"gochat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindMesages(c *gin.Context) {
	var messages []models.Message
	models.DB.Find(&messages)

	c.JSON(http.StatusOK, gin.H{
		"data": messages,
	})
}
