package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

type TaxRateRepository interface {
	CreateTaxRate(taxRate models.TaxRate) (models.TaxRate, error)
	GetTaxRateByID(id uint) (models.TaxRate, error)
	UpdateTaxRate(taxRate models.TaxRate) (models.TaxRate, error)
	DeleteTaxRate(id uint) error
	GetTaxRateByState(state string) ([]models.TaxRate, error)
}

type taxRateRepository struct {
	db *gorm.DB
}

func NewTaxRateRepository() TaxRateRepository {
	return &taxRateRepository{db: database.DB}
}

func (r *taxRateRepository) CreateTaxRate(taxRate models.TaxRate) (models.TaxRate, error) {
	err := r.db.Create(&taxRate).Error
	return taxRate, err
}

func (r *taxRateRepository) GetTaxRateByID(id uint) (models.TaxRate, error) {
	var taxRate models.TaxRate
	err := r.db.First(&taxRate, id).Error
	return taxRate, err
}

func (r *taxRateRepository) UpdateTaxRate(taxRate models.TaxRate) (models.TaxRate, error) {
	err := r.db.Save(&taxRate).Error
	return taxRate, err
}

func (r *taxRateRepository) DeleteTaxRate(id uint) error {
	return r.db.Delete(&models.TaxRate{}, id).Error
}

func (r *taxRateRepository) GetTaxRateByState(state string) ([]models.TaxRate, error) {
	var taxRates []models.TaxRate
	err := r.db.Where("state = ?", state).Find(&taxRates).Error
	return taxRates, err
}
