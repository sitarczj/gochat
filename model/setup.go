package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the database!")
	}

	database.AutoMigrate(&Message{})

	DB = database
}
