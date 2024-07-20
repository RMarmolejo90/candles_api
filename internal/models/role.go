package models

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type UserRole struct {
	gorm.Model
	UserID uint
	RoleID uint
}
