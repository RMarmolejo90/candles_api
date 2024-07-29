package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

// Interfaces
type CandleRepository interface {
	CreateCandle(candle models.Candle) (models.Candle, error)
	GetCandleByID(id uint) (models.Candle, error)
	GetCandles() ([]models.Candle, error)
	UpdateCandle(candle models.Candle) (models.Candle, error)
	DeleteCandle(id uint) error
}

type ImageRepository interface {
	CreateImage(image models.Image) (models.Image, error)
	DeleteImage(id uint) error
}

type StockQuantityRepository interface {
	CreateStock(stock models.StockQuantity) (models.StockQuantity, error)
	GetStockByID(id uint) (models.StockQuantity, error)
	UpdateStock(stock models.StockQuantity) (models.StockQuantity, error)
	DeleteStock(id uint) error
}

// Structs for each repository
type candleRepository struct {
	db *gorm.DB
}

type imageRepository struct {
	db *gorm.DB
}

type stockQuantityRepository struct {
	db *gorm.DB
}

// New repository constructors
func NewCandleRepository() CandleRepository {
	return &candleRepository{db: database.DB}
}

func NewImageRepository() ImageRepository {
	return &imageRepository{db: database.DB}
}

func NewStockQuantityRepository() StockQuantityRepository {
	return &stockQuantityRepository{db: database.DB}
}

// Implementations for CandleRepository
func (r *candleRepository) CreateCandle(candle models.Candle) (models.Candle, error) {
	err := r.db.Create(&candle).Error
	return candle, err
}

func (r *candleRepository) GetCandleByID(id uint) (models.Candle, error) {
	var candle models.Candle
	err := r.db.First(&candle, id).Error
	return candle, err
}

func (r *candleRepository) GetCandles() ([]models.Candle, error) {
	var candles []models.Candle
	err := r.db.Find(&candles).Error
	return candles, err
}

func (r *candleRepository) UpdateCandle(candle models.Candle) (models.Candle, error) {
	err := r.db.Save(&candle).Error
	return candle, err
}

func (r *candleRepository) DeleteCandle(id uint) error {
	return r.db.Delete(&models.Candle{}, id).Error
}

// Implementations for ImageRepository
func (r *imageRepository) CreateImage(image models.Image) (models.Image, error) {
	err := r.db.Create(&image).Error
	return image, err
}

func (r *imageRepository) DeleteImage(id uint) error {
	var image models.Image
	if err := r.db.First(&image, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&image, id).Error
}

// Implementations for StockQuantityRepository
func (r *stockQuantityRepository) CreateStock(stock models.StockQuantity) (models.StockQuantity, error) {
	err := r.db.Create(&stock).Error
	return stock, err
}

func (r *stockQuantityRepository) GetStockByID(id uint) (models.StockQuantity, error) {
	var stock models.StockQuantity
	err := r.db.First(&stock, id).Error
	return stock, err
}

func (r *stockQuantityRepository) UpdateStock(stock models.StockQuantity) (models.StockQuantity, error) {
	err := r.db.Save(&stock).Error
	return stock, err
}

func (r *stockQuantityRepository) DeleteStock(id uint) error {
	return r.db.Delete(&models.StockQuantity{}, id).Error
}
