package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

// RoleRepository defines the methods for interacting with Role data
type RoleRepository interface {
	CreateRole(role models.Role) (models.Role, error)
	GetRoleByID(id uint) (models.Role, error)
	GetRoleByName(name string) (models.Role, error)
	UpdateRole(role models.Role) (models.Role, error)
	DeleteRole(id uint) error
}

// UserRoleRepository defines the methods for interacting with UserRole data
type UserRoleRepository interface {
	CreateUserRole(userRole models.UserRole) (models.UserRole, error)
	GetUserRoleByID(id uint) (models.UserRole, error)
	GetRolesByUserID(userID uint) ([]models.Role, error)
	GetUsersByRoleID(roleID uint) ([]uint, error) // Returns a list of user IDs
	DeleteUserRole(id uint) error
}

type roleRepository struct {
	db *gorm.DB
}

type userRoleRepository struct {
	db *gorm.DB
}

// NewRoleRepository creates a new instance of RoleRepository
func NewRoleRepository() RoleRepository {
	return &roleRepository{db: database.DB}
}

// NewUserRoleRepository creates a new instance of UserRoleRepository
func NewUserRoleRepository() UserRoleRepository {
	return &userRoleRepository{db: database.DB}
}

// RoleRepository implementation
func (r *roleRepository) CreateRole(role models.Role) (models.Role, error) {
	err := r.db.Create(&role).Error
	return role, err
}

func (r *roleRepository) GetRoleByID(id uint) (models.Role, error) {
	var role models.Role
	err := r.db.First(&role, id).Error
	return role, err
}

func (r *roleRepository) GetRoleByName(name string) (models.Role, error) {
	var role models.Role
	err := r.db.Where("name = ?", name).First(&role).Error
	return role, err
}

func (r *roleRepository) UpdateRole(role models.Role) (models.Role, error) {
	err := r.db.Save(&role).Error
	return role, err
}

func (r *roleRepository) DeleteRole(id uint) error {
	return r.db.Delete(&models.Role{}, id).Error
}

// UserRoleRepository implementation
func (r *userRoleRepository) CreateUserRole(userRole models.UserRole) (models.UserRole, error) {
	err := r.db.Create(&userRole).Error
	return userRole, err
}

func (r *userRoleRepository) GetUserRoleByID(id uint) (models.UserRole, error) {
	var userRole models.UserRole
	err := r.db.First(&userRole, id).Error
	return userRole, err
}

func (r *userRoleRepository) GetRolesByUserID(userID uint) ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Joins("JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}

func (r *userRoleRepository) GetUsersByRoleID(roleID uint) ([]uint, error) {
	var userRoles []models.UserRole
	err := r.db.Where("role_id = ?", roleID).Find(&userRoles).Error

	if err != nil {
		return nil, err
	}

	userIDs := make([]uint, len(userRoles))
	for i, userRole := range userRoles {
		userIDs[i] = userRole.UserID
	}

	return userIDs, nil
}

func (r *userRoleRepository) DeleteUserRole(id uint) error {
	return r.db.Delete(&models.UserRole{}, id).Error
}
