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
	// Business logic or validation can be added here before saving
	// Example: Ensure the address has all required fields
	if address.City == "" || address.State == "" || address.PostalCode == "" || address.Country == "" {
		return models.Address{}, errors.New("all address fields must be provided")
	}
	return s.addressRepo.CreateAddress(address)
}

func (s *addressService) GetAddressByID(id uint) (models.Address, error) {
	return s.addressRepo.GetAddressByID(id)
}

func (s *addressService) UpdateAddress(id uint, input models.Address) (models.Address, error) {
	// Retrieve existing address, if necessary
	existingAddress, err := s.addressRepo.GetAddressByID(id)
	if err != nil {
		return models.Address{}, err
	}

	// Update fields, ensuring that specific fields are maintained
	existingAddress.City = input.City
	existingAddress.State = input.State
	existingAddress.PostalCode = input.PostalCode
	existingAddress.Country = input.Country

	return s.addressRepo.UpdateAddress(id, existingAddress)
}

func (s *addressService) DeleteAddress(id uint) error {
	// Additional business logic before deletion can be added here
	return s.addressRepo.DeleteAddress(id)
}

func (s *addressService) GetAddressesByUserID(userID uint) ([]models.Address, error) {
	// This method can be added to the repository if needed
	var addresses []models.Address
	err := s.addressRepo.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}
