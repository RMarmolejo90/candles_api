package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(review models.Review) (models.Review, error)
	GetReviewByID(id uint) (models.Review, error)
	UpdateReview(review models.Review) (models.Review, error)
	DeleteReview(id uint) error
	GetReviewsByUserID(userID uint) ([]models.Review, error)
	GetReviewsByCandleID(candleID uint) ([]models.Review, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository() ReviewRepository {
	return &reviewRepository{db: database.DB}
}

func (r *reviewRepository) CreateReview(review models.Review) (models.Review, error) {
	err := r.db.Create(&review).Error
	return review, err
}

func (r *reviewRepository) GetReviewByID(id uint) (models.Review, error) {
	var review models.Review
	err := r.db.First(&review, id).Error
	return review, err
}

func (r *reviewRepository) UpdateReview(review models.Review) (models.Review, error) {
	err := r.db.Save(&review).Error
	return review, err
}

func (r *reviewRepository) DeleteReview(id uint) error {
	return r.db.Delete(&models.Review{}, id).Error
}

func (r *reviewRepository) GetReviewsByUserID(userID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("user_id = ?", userID).Find(&reviews).Error
	return reviews, err
}

func (r *reviewRepository) GetReviewsByCandleID(candleID uint) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("candle_id = ?", candleID).Find(&reviews).Error
	return reviews, err
}
