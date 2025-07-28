package models

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Name    string        `gorm:"not null"`
	Members []ChatMember  `gorm:"foreignKey:ChatID"`
}