package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `gorm:"unique;not null"`
	Email     string     `gorm:"unique;not null"`
	Password  string     `gorm:"not null"`
	Addresses []Address  `gorm:"foreignKey:UserID"`
	Cart      Cart       `gorm:"foreignKey:UserID"`
	Orders    []Order    `gorm:"foreignKey:UserID"`
	Reviews   []Review   `gorm:"foreignKey:UserID"`
	Roles     []UserRole `gorm:"foreignKey:UserID"`
}
