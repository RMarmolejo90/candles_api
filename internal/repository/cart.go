package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

// CartRepository defines the methods for interacting with the Cart data
type CartRepository interface {
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	GetCartByID(id uint) (models.Cart, error)
	DeleteCart(id uint) error
}

// cartRepository is the concrete implementation of CartRepository
type cartRepository struct {
	db *gorm.DB
}

// NewCartRepository creates a new CartRepository
func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) CreateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Create(&cart).Error
	return cart, err
}

func (r *cartRepository) UpdateCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Save(&cart).Error
	return cart, err
}

func (r *cartRepository) GetCartByID(id uint) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("CartItems").First(&cart, id).Error
	return cart, err
}

func (r *cartRepository) DeleteCart(id uint) error {
	return r.db.Delete(&models.Cart{}, id).Error
}
