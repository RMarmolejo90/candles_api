package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
)

func CreateCart(input models.Cart) (models.Cart, error) {
	var cart models.Cart
	err := database.DB.Create(&input).Error
	return cart, err
}

func UpdateCart(id string, input models.Cart) (models.Cart, error) {
	var cart models.Cart

	if err := database.DB.First(&cart, id).Error; err != nil {
		return cart, err
	}
	err := database.DB.Model(&cart).Updates(&input).Error
	return cart, err
}
