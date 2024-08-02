package services

import (
	"errors"
	"time"

	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type CouponService interface {
	CreateCoupon(coupon models.Coupon) (models.Coupon, error)
	GetCouponByID(id uint) (models.Coupon, error)
	GetCouponByCode(code string) (models.Coupon, error)
	UpdateCoupon(coupon models.Coupon) (models.Coupon, error)
	DeleteCoupon(id uint) error
	GetValidCoupons() ([]models.Coupon, error)
}

// Coupons will need encryption and validation added before implementing
// this is just here to expand on in the future
// there are no current plans to use coupons

type couponService struct {
	couponRepo repository.CouponRepository
}

func NewCouponService(couponRepo repository.CouponRepository) CouponService {
	return &couponService{couponRepo: couponRepo}
}

func (s *couponService) CreateCoupon(coupon models.Coupon) (models.Coupon, error) {
	// validation can be added here before saving
	if coupon.Code == "" || coupon.Discount <= 0 || coupon.ValidUntil.Before(time.Now()) {
		return models.Coupon{}, errors.New("invalid coupon details")
	}
	return s.couponRepo.CreateCoupon(coupon)
}

func (s *couponService) GetCouponByID(id uint) (models.Coupon, error) {
	return s.couponRepo.GetCouponByID(id)
}

func (s *couponService) GetCouponByCode(code string) (models.Coupon, error) {
	return s.couponRepo.GetCouponByCode(code)
}

func (s *couponService) UpdateCoupon(coupon models.Coupon) (models.Coupon, error) {

	if coupon.Code == "" || coupon.Discount <= 0 || coupon.ValidUntil.Before(time.Now()) {
		return models.Coupon{}, errors.New("invalid coupon details")
	}
	return s.couponRepo.UpdateCoupon(coupon)
}

func (s *couponService) DeleteCoupon(id uint) error {
	return s.couponRepo.DeleteCoupon(id)
}

func (s *couponService) GetValidCoupons() ([]models.Coupon, error) {
	return s.couponRepo.GetValidCoupons()
}
