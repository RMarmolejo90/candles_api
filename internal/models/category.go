package models

import (
    "gorm.io/gorm"
)

type Category struct {
    gorm.Model
    Name      string    `gorm:"unique;not null"`
}

type CandleCategory struct {
    gorm.Model
    CandleID  uint
    CategoryID uint
}
