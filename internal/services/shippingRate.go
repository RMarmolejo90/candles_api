package services

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type ShippingRateService interface {
	GetShippingRate(zipCode string, quantity int) (float64, error)
}

type shippingRateService struct {
	shippingRateRepo repository.ShippingRateRepository
	cache            *cache.Cache
	uspsAPIKey       string
}

func NewShippingRateService(shippingRateRepo repository.ShippingRateRepository, uspsAPIKey string) ShippingRateService {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &shippingRateService{
		shippingRateRepo: shippingRateRepo,
		cache:            c,
		uspsAPIKey:       uspsAPIKey,
	}
}

func (s *shippingRateService) GetShippingRate(zipCode string, quantity int) (float64, error) {
	cacheKey := fmt.Sprintf("%s-%d", zipCode, quantity)
	if cachedRate, found := s.cache.Get(cacheKey); found {
		return cachedRate.(float64), nil
	}

	rate, err := s.fetchUSPSRate(zipCode, quantity)
	if err != nil {
		return 0, err
	}

	s.cache.Set(cacheKey, rate, cache.DefaultExpiration)

	shippingRate := models.ShippingRate{
		Region: zipCode,
		Rate:   rate,
	}

	_, dbErr := s.shippingRateRepo.CreateShippingRate(shippingRate)
	if dbErr != nil {
		return 0, dbErr
	}

	return rate, nil
}

func (s *shippingRateService) fetchUSPSRate(zipCode string, quantity int) (string, error) {
	// This function should be implemented to call the USPS API.

	// call api url
	// Make HTTP GET request to USPS API
	// Parse the response and extract the shipping rate

	// this is a placeholder
	// change to a float type and return the actual shipping rate
	rate := "Api call needs to be implemented"

	return rate, nil
}
