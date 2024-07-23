package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
)

func CreateCandle(candle models.Candle) (models.Candle, error) {
	result := database.DB.Create(&candle)
	return candle, result.Error
}

func GetCandleById(id string) (candle models.Candle, error error) {
	result := database.DB.First(&id)
	return candle, result.Error
}

func GetCandles() (candles []models.Candle, error error) {
	result := database.DB.Find(&candles)
	return candles, result.Error
}

func UpdateCandle(id string, input models.Candle) (candle models.Candle, error error) {
	if err := database.DB.First(&candle, id); err != nil {
		return candle, err.Error
	}
	result := database.DB.Model(&candle).Updates(&input)
	return candle, result.Error
}

func DeleteCandle(id string) error {
	var candle models.Candle
	if err := database.DB.First(&candle); err != nil {
		return err.Error
	}
	result := database.DB.Delete(&candle)
	return result.Error
}
