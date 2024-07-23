package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
)

func CreateCart(input models.Cart) (cart models.Cart, error error) {
	result := database.DB.Create(&input)
	return cart, result.Error
}

func UpdateCart(id string, input models.Cart) (cart models.Cart, error error) {
	if err := database.DB.First(&cart, id); err != nil {
		return cart, err.Error
	}
	result := database.DB.Model(&cart).Updates(&input)
	return cart, result.Error
}
