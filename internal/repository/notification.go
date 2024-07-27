package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(notification models.Notification) (models.Notification, error)
	GetNotificationByID(id uint) (models.Notification, error)
	GetNotificationsByUserID(userID uint) ([]models.Notification, error)
	MarkAsRead(notificationID uint) error
	DeleteNotification(id uint) error
	GetUnreadNotifications(userID uint) ([]models.Notification, error)
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepository{db: database.DB}
}

func (r *notificationRepository) CreateNotification(notification models.Notification) (models.Notification, error) {
	err := r.db.Create(&notification).Error
	return notification, err
}

func (r *notificationRepository) GetNotificationByID(id uint) (models.Notification, error) {
	var notification models.Notification
	err := r.db.First(&notification, id).Error
	return notification, err
}

func (r *notificationRepository) GetNotificationsByUserID(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("user_id = ?", userID).Find(&notifications).Error
	return notifications, err
}

func (r *notificationRepository) MarkAsRead(notificationID uint) error {
	return r.db.Model(&models.Notification{}).Where("id = ?", notificationID).Update("read", true).Error
}

func (r *notificationRepository) DeleteNotification(id uint) error {
	return r.db.Delete(&models.Notification{}, id).Error
}

func (r *notificationRepository) GetUnreadNotifications(userID uint) ([]models.Notification, error) {
	var notifications []models.Notification
	err := r.db.Where("user_id = ? AND read = ?", userID, false).Find(&notifications).Error
	return notifications, err
}
