package services

import (
	"errors"

	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type AddressService interface {
	CreateAddress(address models.Address) (models.Address, error)
	GetAddressByID(id uint) (models.Address, error)
	UpdateAddress(id uint, input models.Address) (models.Address, error)
	DeleteAddress(id uint) error
	GetAddressesByUserID(userID uint) ([]models.Address, error)
}

type addressService struct {
	addressRepo repository.AddressRepository
}

func NewAddressService(addressRepo repository.AddressRepository) AddressService {
	return &addressService{addressRepo: addressRepo}
}

func (s *addressService) CreateAddress(address models.Address) (models.Address, error) {
	if address.City == "" || address.State == "" || address.PostalCode == "" || address.Country == "" {
		return models.Address{}, errors.New("all address fields must be provided")
	}
	return s.addressRepo.CreateAddress(address)
}

func (s *addressService) GetAddressByID(id uint) (models.Address, error) {
	return s.addressRepo.GetAddressByID(id)
}

func (s *addressService) GetAddressesByUserID(userID uint) ([]models.Address, error) {
	return s.addressRepo.GetAddressesByUserID(userID)
}

func (s *addressService) UpdateAddress(id uint, input models.Address) (models.Address, error) {
	return s.addressRepo.UpdateAddress(id, input)
}

func (s *addressService) DeleteAddress(id uint) error {
	return s.addressRepo.DeleteAddress(id)
}
