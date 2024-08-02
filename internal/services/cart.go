package services

import (
	"errors"

	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type CartService interface {
	CreateCart(cart models.Cart) (models.Cart, error)
	UpdateCart(cart models.Cart) (models.Cart, error)
	GetCartByID(id uint) (models.Cart, error)
	DeleteCart(id uint) error
	AddItemToCart(cartID uint, item models.CartItem) (models.Cart, error)
	RemoveItemFromCart(cartID, itemID uint) (models.Cart, error)
}

type cartService struct {
	cartRepo repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) CartService {
	return &cartService{cartRepo: cartRepo}
}

func (s *cartService) CreateCart(cart models.Cart) (models.Cart, error) {
	// need to check if the user already has a cart before creating a new one

	if cart.UserID == 0 {
		return models.Cart{}, errors.New("user ID must be provided")
	}
	return s.cartRepo.CreateCart(cart)
}

func (s *cartService) UpdateCart(cart models.Cart) (models.Cart, error) {
	return s.cartRepo.UpdateCart(cart)
}

func (s *cartService) GetCartByID(id uint) (models.Cart, error) {
	return s.cartRepo.GetCartByID(id)
}

func (s *cartService) DeleteCart(id uint) error {
	return s.cartRepo.DeleteCart(id)
}

func (s *cartService) AddItemToCart(cartID uint, item models.CartItem) (models.Cart, error) {
	cart, err := s.cartRepo.GetCartByID(cartID)
	if err != nil {
		return models.Cart{}, err
	}
	// add item to cart
	cart.CartItems = append(cart.CartItems, item)
	return s.cartRepo.UpdateCart(cart)
}

func (s *cartService) RemoveItemFromCart(cartID, itemID uint) (models.Cart, error) {
	cart, err := s.cartRepo.GetCartByID(cartID)
	if err != nil {
		return models.Cart{}, err
	}

	// Remove item logic
	for i, item := range cart.CartItems {
		if item.ID == itemID {
			cart.CartItems = append(cart.CartItems[:i], cart.CartItems[i+1:]...)
			break
		}
	}
	return s.cartRepo.UpdateCart(cart)
}
