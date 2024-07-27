package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

// CartItemRepository defines the methods for interacting with the CartItem data
type CartItemRepository interface {
	CreateCartItem(cartItem models.CartItem) (models.CartItem, error)
	UpdateCartItem(cartItem models.CartItem) (models.CartItem, error)
	GetCartItemByID(id uint) (models.CartItem, error)
	GetCartItemsByCartID(cartID uint) ([]models.CartItem, error)
	DeleteCartItem(id uint) error
}

// cartItemRepository is the concrete implementation of CartItemRepository
type cartItemRepository struct {
	db *gorm.DB
}

// NewCartItemRepository creates a new CartItemRepository
func NewCartItemRepository() CartItemRepository {
	return &cartItemRepository{db: database.DB}
}

func (r *cartItemRepository) CreateCartItem(cartItem models.CartItem) (models.CartItem, error) {
	err := r.db.Create(&cartItem).Error
	return cartItem, err
}

func (r *cartItemRepository) UpdateCartItem(cartItem models.CartItem) (models.CartItem, error) {
	err := r.db.Save(&cartItem).Error
	return cartItem, err
}

func (r *cartItemRepository) GetCartItemByID(id uint) (models.CartItem, error) {
	var cartItem models.CartItem
	err := r.db.First(&cartItem, id).Error
	return cartItem, err
}

func (r *cartItemRepository) GetCartItemsByCartID(cartID uint) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	err := r.db.Where("cart_id = ?", cartID).Find(&cartItems).Error
	return cartItems, err
}

func (r *cartItemRepository) DeleteCartItem(id uint) error {
	return r.db.Delete(&models.CartItem{}, id).Error
}
