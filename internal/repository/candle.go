package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
)

func CreateCandle(candle models.Candle) (models.Candle, error) {
	err := database.DB.Create(&candle).Error
	return candle, err
}

func GetCandleById(id string) (models.Candle, error) {
	var candle models.Candle
	err := database.DB.First(&candle, id).Error
	return candle, err
}

func GetCandles() ([]models.Candle, error) {
	var candles []models.Candle
	err := database.DB.Find(&candles).Error
	return candles, err
}

func UpdateCandle(id string, input models.Candle) (models.Candle, error) {
	var candle models.Candle
	if err := database.DB.First(&candle, id).Error; err != nil {
		return candle, err
	}
	err := database.DB.Model(&candle).Updates(&input).Error
	return candle, err
}

func DeleteCandle(id string) error {
	var candle models.Candle
	if err := database.DB.First(&candle, id).Error; err != nil {
		return err
	}

	return database.DB.Delete(&candle).Error

}
