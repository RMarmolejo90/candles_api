package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Code       string    `gorm:"unique;not null"`
	Discount   float64   `gorm:"not null"`
	ValidUntil time.Time `gorm:"not null"`
}
