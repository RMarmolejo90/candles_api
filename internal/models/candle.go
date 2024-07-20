package models

import (
	"gorm.io/gorm"
)

type Candle struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Description   string
	Price         float64          `gorm:"not null"`
	Status        string           `gorm:"not null;default:'active'"`   // 'active' or 'deactivated'
	StockStatus   string           `gorm:"not null;default:'in-stock'"` // 'in-stock' or 'sold-out'
	StockQuantity int              `gorm:"not null;default:0"`
	Images        []Image          `gorm:"foreignKey:CandleID"`
	Categories    []CandleCategory `gorm:"foreignKey:CandleID"`
	Reviews       []Review         `gorm:"foreignKey:CandleID"`
}

type Image struct {
	gorm.Model
	CandleID       uint
	URL            string `gorm:"not null"`
	IsDisplayPhoto bool   `gorm:"not null;default:false"`
}

type StockQuantity struct {
	gorm.Model
	CandleID uint
	Quantity int `gorm:"not null"`
}
