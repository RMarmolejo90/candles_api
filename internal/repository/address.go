package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
)

func CreateAddress() (address models.Address, error error) {
	result := database.DB.Create(&address)
	return address, result.Error
}

func GetAddressByID(id string) (address models.Address, error error) {
	result := database.DB.First(&id)
	return address, result.Error
}

func UpdateAddress(id string, input models.Address) (address models.Address, error error) {
	if err := database.DB.First(&address, id); err != nil {
		return address, err.Error
	}

	result := database.DB.

}
