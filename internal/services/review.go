package services

import (
	"errors"

	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type ReviewService interface {
	CreateReview(review models.Review) (models.Review, error)
	GetReviewByID(id uint) (models.Review, error)
	UpdateReview(review models.Review) (models.Review, error)
	DeleteReview(id uint) error
	GetReviewsByUserID(userID uint) ([]models.Review, error)
	GetReviewsByCandleID(candleID uint) ([]models.Review, error)
}

type reviewService struct {
	reviewRepo repository.ReviewRepository
}

func NewReviewService(reviewRepo repository.ReviewRepository) ReviewService {
	return &reviewService{reviewRepo: reviewRepo}
}

func (s *reviewService) CreateReview(review models.Review) (models.Review, error) {
	if review.UserID == 0 || review.CandleID == 0 || review.Rating == 0 {
		return models.Review{}, errors.New("all review fields must be provided")
	}
	return s.reviewRepo.CreateReview(review)
}

func (s *reviewService) GetReviewByID(id uint) (models.Review, error) {
	return s.reviewRepo.GetReviewByID(id)
}

func (s *reviewService) UpdateReview(review models.Review) (models.Review, error) {
	if review.UserID == 0 || review.CandleID == 0 || review.Rating == 0 {
		return models.Review{}, errors.New("all review fields must be provided")
	}
	return s.reviewRepo.UpdateReview(review)
}

func (s *reviewService) DeleteReview(id uint) error {
	return s.reviewRepo.DeleteReview(id)
}

func (s *reviewService) GetReviewsByUserID(userID uint) ([]models.Review, error) {
	return s.reviewRepo.GetReviewsByUserID(userID)
}

func (s *reviewService) GetReviewsByCandleID(candleID uint) ([]models.Review, error) {
	return s.reviewRepo.GetReviewsByCandleID(candleID)
}
