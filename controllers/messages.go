package controllers

import (
	"gochat/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMessageInput struct {
	Content string `json:"content" binding:"required"`
}

func FindMesages(c *gin.Context) {
	var messages []models.Message
	models.DB.Find(&messages)

	c.JSON(http.StatusOK, gin.H{
		"data": messages,
	})
}

func CreateMessage(c *gin.Context) {
	var input CreateMessageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	message := models.Message{
		Content: input.Content,
	}
	models.DB.Create(&message)

	c.JSON(http.StatusCreated, gin.H{
		"data": message,
	})
}
