package main

import (
	"gochat/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectToDatabase()

	r.Run()
}
