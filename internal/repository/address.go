package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

type AddressRepository interface {
	CreateAddress(address models.Address) (models.Address, error)
	GetAddressByID(id uint) (models.Address, error)
	UpdateAddress(id uint, input models.Address) (models.Address, error)
	DeleteAddress(id uint) error
}

type addressRepository struct {
	db *gorm.DB
}

func NewAddressRepository() AddressRepository {
	return &addressRepository{db: database.DB}
}

// Ensure the method receiver is *addressRepository, not the package-level database.DB
func (r *addressRepository) CreateAddress(address models.Address) (models.Address, error) {
	err := r.db.Create(&address).Error
	return address, err
}

func (r *addressRepository) GetAddressByID(id uint) (models.Address, error) {
	var address models.Address
	err := r.db.First(&address, id).Error
	return address, err
}

func (r *addressRepository) UpdateAddress(id uint, input models.Address) (models.Address, error) {
	var address models.Address
	if err := r.db.First(&address, id).Error; err != nil {
		return address, err
	}

	err := r.db.Model(&address).Updates(input).Error
	return address, err
}

func (r *addressRepository) DeleteAddress(id uint) error {
	return r.db.Delete(&models.Address{}, id).Error
}
