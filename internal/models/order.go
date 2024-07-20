package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID          uint
	AddressID       uint
	TotalAmount     float64        `gorm:"not null"`
	TaxAmount       float64        `gorm:"not null"`
	ShippingAmount  float64        `gorm:"not null"`
	OrderStatus     string         `gorm:"not null;default:'ordered'"` // 'ordered', 'in process', 'shipped', 'complete'
	StripePaymentID string         `gorm:"not null"`                   // To store Stripe Payment ID
	OrderItems      []OrderItem    `gorm:"foreignKey:OrderID"`
	OrderHistories  []OrderHistory `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID  uint
	CandleID uint
	Quantity int     `gorm:"not null"`
	Price    float64 `gorm:"not null"`
}

type OrderStatus struct {
	gorm.Model
	StatusName string `gorm:"not null"` // 'ordered', 'in process', 'shipped', 'complete'
}

type OrderHistory struct {
	gorm.Model
	OrderID   uint
	Status    string    `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}
