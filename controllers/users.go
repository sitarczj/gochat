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

type LoginUserInput struct {
	Username      string `json:"username" binding:"required"`
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

	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

func Login(c *gin.Context) {
	var input LoginUserInput
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.DB.Where("Username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})

		return
	}

	if !services.IsPasswordValid(input.PlainPassword, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})

		return
	}

	authToken := services.GenerateToken()

	token := models.Token{
		UserID: user.ID,
		Token:  authToken,
	}

	models.DB.Create(&token)

	c.JSON(http.StatusOK, gin.H{
		"token": authToken,
	})
}
