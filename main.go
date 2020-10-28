package main

import (
	"gochat/controllers"
	"gochat/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectToDatabase()

	r.GET("/messages", controllers.FindMesages)
	r.POST("/messages", controllers.CreateMessage)

	r.Run()
}
