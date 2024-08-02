package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type TaxRateService interface {
	GetTaxRate(zipCode string) (float64, error)
}

type taxRateService struct {
	taxRateRepo repository.TaxRateRepository
	cache       *cache.Cache
	taxAPIKey   string
}

func NewTaxRateService(taxRateRepo repository.TaxRateRepository, taxAPIKey string) TaxRateService {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &taxRateService{
		taxRateRepo: taxRateRepo,
		cache:       c,
		taxAPIKey:   taxAPIKey,
	}
}

func (s *taxRateService) GetTaxRate(zipCode string) (float64, error) {
	cacheKey := fmt.Sprintf("%s", zipCode)
	if cachedRate, found := s.cache.Get(cacheKey); found {
		return cachedRate.(float64), nil
	}

	rate, err := s.fetchTaxRate(zipCode)
	if err != nil {
		return 0, err
	}

	s.cache.Set(cacheKey, rate, cache.DefaultExpiration)

	taxRate := models.TaxRate{
		State: zipCode,
		Rate:  rate,
	}

	_, dbErr := s.taxRateRepo.CreateTaxRate(taxRate)
	if dbErr != nil {
		return 0, dbErr
	}

	return rate, nil
}

func (s *taxRateService) fetchTaxRate(zipCode string) (float64, error) {
	// Example of making an HTTP GET request to an external API
	apiURL := fmt.Sprintf("https://api.example.com/taxrate?zipCode=%s&apiKey=%s", zipCode, s.taxAPIKey)
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, errors.New("failed to fetch rate from tax API")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result struct {
		Rate float64 `json:"rate"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}

	return result.Rate, nil
}
