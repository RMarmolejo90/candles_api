package services

import (
	"errors"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type UserService interface {
	CreateUser(user models.User) (models.User, error)
	GetUserByID(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return models.User{}, errors.New("username, email, and password must be provided")
	}
	return s.userRepo.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) GetUserByEmail(email string) (models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetUserByUsername(username string) (models.User, error) {
	return s.userRepo.GetUserByUsername(username)
}

func (s *userService) UpdateUser(user models.User) (models.User, error) {
	if user.Username == "" || user.Email == "" {
		return models.User{}, errors.New("username and email must be provided")
	}
	return s.userRepo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}
