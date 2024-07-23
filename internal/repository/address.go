package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
)

func CreateAddress(address models.Address) (models.Address, error) {
	result := database.DB.Create(&address)
	return address, result.Error
}

func GetAddressByID(id string) (models.Address, error) {
	var address models.Address

	result := database.DB.First(&address, id)
	return address, result.Error
}

func UpdateAddress(id string, input models.Address) (models.Address, error) {
	var address models.Address

	if err := database.DB.First(&address, id).Error; err != nil {
		return address, err
	}

	if err := database.DB.Model(&address).Updates(input).Error; err != nil {
		return address, err
	}

	return address, nil

}

func DeleteAddress(id string) error {
	var address models.Address
	if err := database.DB.First(&address, id).Error; err != nil {
		return err
	}
	return database.DB.Delete(&address).Error

}
