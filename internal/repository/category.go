package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
)

func CreateCategory(category models.Category) (models.Category, error) {
	err := database.DB.Create(&category).Error
	return category, err
}

func UpdateCategory(id string, input models.Category) (models.Category, error) {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return category, err
	}

	err := database.DB.Model(&category).Updates(input).Error
	return category, err
}

func GetCategory(id string) (models.Category, error) {
	var category models.Category
	err := database.DB.First(&category, id).Error
	return category, err
}

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&category).Error
	return []categories, err
}
