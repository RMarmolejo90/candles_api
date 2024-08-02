package services

import (
	"errors"

	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

type RoleService interface {
	CreateRole(role models.Role) (models.Role, error)
	GetRoleByID(id uint) (models.Role, error)
	GetRoleByName(name string) (models.Role, error)
	UpdateRole(role models.Role) (models.Role, error)
	DeleteRole(id uint) error
}

type UserRoleService interface {
	CreateUserRole(userRole models.UserRole) (models.UserRole, error)
	GetUserRoleByID(id uint) (models.UserRole, error)
	GetRolesByUserID(userID uint) ([]models.Role, error)
	GetUsersByRoleID(roleID uint) ([]uint, error)
	DeleteUserRole(id uint) error
}

type roleService struct {
	roleRepo repository.RoleRepository
}

type userRoleService struct {
	userRoleRepo repository.UserRoleRepository
}

func NewRoleService(roleRepo repository.RoleRepository) RoleService {
	return &roleService{roleRepo: roleRepo}
}

func NewUserRoleService(userRoleRepo repository.UserRoleRepository) UserRoleService {
	return &userRoleService{userRoleRepo: userRoleRepo}
}

// RoleService implementation
func (s *roleService) CreateRole(role models.Role) (models.Role, error) {
	if role.Name == "" {
		return models.Role{}, errors.New("role name must be provided")
	}
	return s.roleRepo.CreateRole(role)
}

func (s *roleService) GetRoleByID(id uint) (models.Role, error) {
	return s.roleRepo.GetRoleByID(id)
}

func (s *roleService) GetRoleByName(name string) (models.Role, error) {
	return s.roleRepo.GetRoleByName(name)
}

func (s *roleService) UpdateRole(role models.Role) (models.Role, error) {
	if role.Name == "" {
		return models.Role{}, errors.New("role name must be provided")
	}
	return s.roleRepo.UpdateRole(role)
}

func (s *roleService) DeleteRole(id uint) error {
	return s.roleRepo.DeleteRole(id)
}

// UserRoleService implementation
func (s *userRoleService) CreateUserRole(userRole models.UserRole) (models.UserRole, error) {
	if userRole.UserID == 0 || userRole.RoleID == 0 {
		return models.UserRole{}, errors.New("user ID and role ID must be provided")
	}
	return s.userRoleRepo.CreateUserRole(userRole)
}

func (s *userRoleService) GetUserRoleByID(id uint) (models.UserRole, error) {
	return s.userRoleRepo.GetUserRoleByID(id)
}

func (s *userRoleService) GetRolesByUserID(userID uint) ([]models.Role, error) {
	return s.userRoleRepo.GetRolesByUserID(userID)
}

func (s *userRoleService) GetUsersByRoleID(roleID uint) ([]uint, error) {
	return s.userRoleRepo.GetUsersByRoleID(roleID)
}

func (s *userRoleService) DeleteUserRole(id uint) error {
	return s.userRoleRepo.DeleteUserRole(id)
}
