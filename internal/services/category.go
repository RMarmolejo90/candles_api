package services

import (
	"errors"

	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type CategoryService interface {
	CreateCategory(category models.Category) (models.Category, error)
	GetCategory(id uint) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	GetAllCategories() ([]models.Category, error)
	DeleteCategory(id uint) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) CreateCategory(category models.Category) (models.Category, error) {
	// validation can be added here before saving
	if category.Name == "" {
		return models.Category{}, errors.New("category name must be provided")
	}
	return s.categoryRepo.CreateCategory(category)
}

func (s *categoryService) GetCategory(id uint) (models.Category, error) {
	return s.categoryRepo.GetCategory(id)
}

func (s *categoryService) UpdateCategory(category models.Category) (models.Category, error) {
	if category.Name == "" {
		return models.Category{}, errors.New("category name must be provided")
	}
	return s.categoryRepo.UpdateCategory(category)
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.categoryRepo.GetAllCategories()
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.categoryRepo.DeleteCategory(id)
}
