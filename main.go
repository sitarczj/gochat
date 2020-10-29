package main

import (
	"gochat/controllers"
	"gochat/middleware"
	"gochat/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectToDatabase()

	secured := r.Group("/")

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	secured.Use(middleware.Authorized)
	{
		secured.GET("/messages", controllers.FindMesages)
		secured.POST("/messages", controllers.CreateMessage)
	}

	r.Run()
}
