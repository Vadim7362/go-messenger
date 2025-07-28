package models

import (
	"gorm.io/gorm"
)

type ChatMember struct {
	gorm.Model
	UserID uint
	ChatID uint
}