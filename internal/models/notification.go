package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	UserID  uint
	Message string `gorm:"not null"`
	Read    bool   `gorm:"not null;default:false"`
}
