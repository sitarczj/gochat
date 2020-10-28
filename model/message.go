package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primary_key"`
	Content string `json:"content"`
}
