package models

import (
	"gorm.io/gorm"
)

type ShippingRate struct {
	gorm.Model
	Region string  `gorm:"not null"`
	Rate   float64 `gorm:"not null"`
}
