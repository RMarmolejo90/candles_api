package models

import (
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartID   uint
	CandleID uint
	Quantity int `gorm:"not null"`
}

type Cart struct {
	gorm.Model
	UserID    uint       `gorm:"not null;unique"` // Ensure each user has only one cart
	CartItems []CartItem `gorm:"foreignKey:CartID"`
}
