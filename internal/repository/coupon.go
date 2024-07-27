package repository

import (
	"time"

	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

type CouponRepository interface {
	CreateCoupon(coupon models.Coupon) (models.Coupon, error)
	GetCouponByID(id uint) (models.Coupon, error)
	GetCouponByCode(code string) (models.Coupon, error)
	UpdateCoupon(coupon models.Coupon) (models.Coupon, error)
	DeleteCoupon(id uint) error
	GetValidCoupons() ([]models.Coupon, error)
}

type couponRepository struct {
	db *gorm.DB
}

func NewCouponRepository() CouponRepository {
	return &couponRepository{db: database.DB}
}

func (r *couponRepository) CreateCoupon(coupon models.Coupon) (models.Coupon, error) {
	err := r.db.Create(&coupon).Error
	return coupon, err
}

func (r *couponRepository) GetCouponByID(id uint) (models.Coupon, error) {
	var coupon models.Coupon
	err := r.db.First(&coupon, id).Error
	return coupon, err
}

func (r *couponRepository) GetCouponByCode(code string) (models.Coupon, error) {
	var coupon models.Coupon
	err := r.db.Where("code = ?", code).First(&coupon).Error
	return coupon, err
}

func (r *couponRepository) UpdateCoupon(coupon models.Coupon) (models.Coupon, error) {
	err := r.db.Save(&coupon).Error
	return coupon, err
}

func (r *couponRepository) DeleteCoupon(id uint) error {
	return r.db.Delete(&models.Coupon{}, id).Error
}

func (r *couponRepository) GetValidCoupons() ([]models.Coupon, error) {
	var coupons []models.Coupon
	now := time.Now()
	err := r.db.Where("valid_until >= ?", now).Find(&coupons).Error
	return coupons, err
}
