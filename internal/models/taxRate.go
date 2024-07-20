package models

import (
	"gorm.io/gorm"
)

type TaxRate struct {
	gorm.Model
	State string  `gorm:"not null"`
	Rate  float64 `gorm:"not null"`
}
