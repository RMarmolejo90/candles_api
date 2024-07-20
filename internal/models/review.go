package models

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID   uint
	CandleID uint
	Rating   int `gorm:"not null"`
	Comment  string
}
