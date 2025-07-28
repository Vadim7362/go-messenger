package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Content string `gorm:"not null"`

	SenderID uint
	Sender   User

	ChatID uint
	Chat   Chat
}