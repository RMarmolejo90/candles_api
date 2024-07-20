package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID       uint
	AddressLine1 string `gorm:"not null"`
	AddressLine2 string
	City         string `gorm:"not null"`
	State        string `gorm:"not null"`
	PostalCode   string `gorm:"not null"`
	Country      string `gorm:"not null"`
}
