package services

import (
	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

// CandleService defines the methods for handling candles, images, and stock quantities.
type CandleService interface {
	// Candle operations
	CreateNewCandle(candle models.Candle) (models.Candle, error)
	GetCandleDetails(id uint) (models.Candle, error)
	UpdateCandleInfo(candle models.Candle) (models.Candle, error)
	RemoveCandle(id uint) error

	// Image operations
	AddImageToCandle(image models.Image) (models.Image, error)
	RemoveImage(id uint) error

	// Stock operations
	CreateStock(stock models.StockQuantity) (models.StockQuantity, error)
	GetStockDetails(id uint) (models.StockQuantity, error)
	UpdateStockQuantity(stock models.StockQuantity) (models.StockQuantity, error)
	RemoveStock(id uint) error
}

type candleService struct {
	candleRepo repository.CandleRepository
	imageRepo  repository.ImageRepository
	stockRepo  repository.StockQuantityRepository
}

// NewCandleService creates a new instance of CandleService.
func NewCandleService(candleRepo repository.CandleRepository, imageRepo repository.ImageRepository, stockRepo repository.StockQuantityRepository) CandleService {
	return &candleService{
		candleRepo: candleRepo,
		imageRepo:  imageRepo,
		stockRepo:  stockRepo,
	}
}

// Candle operations

func (s *candleService) CreateNewCandle(candle models.Candle) (models.Candle, error) {
	return s.candleRepo.CreateCandle(candle)
}

func (s *candleService) GetCandleDetails(id uint) (models.Candle, error) {
	return s.candleRepo.GetCandleByID(id)
}

func (s *candleService) UpdateCandleInfo(candle models.Candle) (models.Candle, error) {
	return s.candleRepo.UpdateCandle(candle)
}

func (s *candleService) RemoveCandle(id uint) error {
	return s.candleRepo.DeleteCandle(id)
}

// Image operations

func (s *candleService) AddImageToCandle(image models.Image) (models.Image, error) {
	return s.imageRepo.CreateImage(image)
}

func (s *candleService) RemoveImage(id uint) error {
	return s.imageRepo.DeleteImage(id)
}

// Stock operations

func (s *candleService) CreateStock(stock models.StockQuantity) (models.StockQuantity, error) {
	return s.stockRepo.CreateStock(stock)
}

func (s *candleService) GetStockDetails(id uint) (models.StockQuantity, error) {
	return s.stockRepo.GetStockByID(id)
}

func (s *candleService) UpdateStockQuantity(stock models.StockQuantity) (models.StockQuantity, error) {
	return s.stockRepo.UpdateStock(stock)
}

func (s *candleService) RemoveStock(id uint) error {
	return s.stockRepo.DeleteStock(id)
}
