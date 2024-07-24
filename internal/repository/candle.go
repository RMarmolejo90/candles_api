package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
)

// Candles

func CreateCandle(candle models.Candle) (models.Candle, error) {
	err := database.DB.Create(&candle).Error
	return candle, err
}

func GetCandleByID(id uint) (models.Candle, error) { // Change to uint for ID type consistency
	var candle models.Candle
	err := database.DB.First(&candle, id).Error
	return candle, err
}

func GetCandles() ([]models.Candle, error) {
	var candles []models.Candle
	err := database.DB.Find(&candles).Error
	return candles, err
}

func UpdateCandle(candle models.Candle) (models.Candle, error) {
	err := database.DB.Save(&candle).Error
	return candle, err
}

func DeleteCandle(id uint) error {
	return database.DB.Delete(&models.Candle{}, id).Error
}

// Images

func CreateImage(image models.Image) (models.Image, error) {
	err := database.DB.Create(&image).Error
	return image, err
}

func DeleteImage(id uint) error {
	var image models.Image
	if err := database.DB.First(&image, id).Error; err != nil {
		return err
	}
	return database.DB.Delete(&image, id).Error
}

// Stock Quantity

func CreateStock(stock models.StockQuantity) (models.StockQuantity, error) {
	err := database.DB.Create(&stock).Error
	return stock, err
}

func GetStockByID(id uint) (models.StockQuantity, error) {
	var stock models.StockQuantity
	err := database.DB.First(&stock, id).Error
	return stock, err
}

func UpdateStock(stock models.StockQuantity) (models.StockQuantity, error) {
	err := database.DB.Save(&stock).Error
	return stock, err
}

func DeleteStock(id uint) error {
	return database.DB.Delete(&models.StockQuantity{}, id).Error
}
