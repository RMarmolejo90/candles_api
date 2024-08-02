package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category models.Category) (models.Category, error)
	GetCategory(id uint) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
	DeleteCategory(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{db: database.DB}
}

func (r *categoryRepository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error
	return category, err
}

func (r *categoryRepository) GetCategory(id uint) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	return category, err
}

func (r *categoryRepository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error
	return category, err
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}
