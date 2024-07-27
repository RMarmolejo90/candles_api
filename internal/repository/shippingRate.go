package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

type ShippingRateRepository interface {
	CreateShippingRate(rate models.ShippingRate) (models.ShippingRate, error)
	GetShippingRateByID(id uint) (models.ShippingRate, error)
	UpdateShippingRate(rate models.ShippingRate) (models.ShippingRate, error)
	DeleteShippingRate(id uint) error
	GetShippingRatesByRegion(region string) ([]models.ShippingRate, error)
}

type shippingRateRepository struct {
	db *gorm.DB
}

func NewShippingRateRepository() ShippingRateRepository {
	return &shippingRateRepository{db: database.DB}
}

func (r *shippingRateRepository) CreateShippingRate(rate models.ShippingRate) (models.ShippingRate, error) {
	err := r.db.Create(&rate).Error
	return rate, err
}

func (r *shippingRateRepository) GetShippingRateByID(id uint) (models.ShippingRate, error) {
	var rate models.ShippingRate
	err := r.db.First(&rate, id).Error
	return rate, err
}

func (r *shippingRateRepository) UpdateShippingRate(rate models.ShippingRate) (models.ShippingRate, error) {
	err := r.db.Save(&rate).Error
	return rate, err
}

func (r *shippingRateRepository) DeleteShippingRate(id uint) error {
	return r.db.Delete(&models.ShippingRate{}, id).Error
}

func (r *shippingRateRepository) GetShippingRatesByRegion(region string) ([]models.ShippingRate, error) {
	var rates []models.ShippingRate
	err := r.db.Where("region = ?", region).Find(&rates).Error
	return rates, err
}
