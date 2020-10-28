package controllers

import (
	"gochat/models"
	"gochat/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterUserInput struct {
	Username      string `json:"username" binding:"required`
	Email         string `json:"email" binding:"required"`
	PlainPassword string `json:"plain_password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	hashedPass, err := services.HasPassword(input.PlainPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something went wrong, please contact adminstrator",
		})

		return
	}

	user := models.User{
		Username: input.Username,
		Password: hashedPass,
		Email:    input.Email,
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
