package repository

import (
	"github.com/rmarmolejo90/candles_api/internal/database"
	"github.com/rmarmolejo90/candles_api/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetOrderByID(id uint) (models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(id uint) error
	GetOrdersByUserID(userID uint) ([]models.Order, error)
}

type OrderItemRepository interface {
	CreateOrderItem(orderItem models.OrderItem) (models.OrderItem, error)
	GetOrderItemByID(id uint) (models.OrderItem, error)
	UpdateOrderItem(orderItem models.OrderItem) (models.OrderItem, error)
	DeleteOrderItem(id uint) error
}

type OrderStatusRepository interface {
	CreateOrderStatus(orderStatus models.OrderStatus) (models.OrderStatus, error)
	GetOrderStatusByID(id uint) (models.OrderStatus, error)
	UpdateOrderStatus(orderStatus models.OrderStatus) (models.OrderStatus, error)
	DeleteOrderStatus(id uint) error
	GetOrderStatusByName(name string) (models.OrderStatus, error)
}

type OrderHistoryRepository interface {
	CreateOrderHistory(orderHistory models.OrderHistory) (models.OrderHistory, error)
	GetOrderHistoryByID(id uint) (models.OrderHistory, error)
	UpdateOrderHistory(orderHistory models.OrderHistory) (models.OrderHistory, error)
	DeleteOrderHistory(id uint) error
	GetOrderHistoriesByOrderID(orderID uint) ([]models.OrderHistory, error)
}

type orderRepository struct {
	db *gorm.DB
}

type orderItemRepository struct {
	db *gorm.DB
}

type orderStatusRepository struct {
	db *gorm.DB
}

type orderHistoryRepository struct {
	db *gorm.DB
}

func NewOrderRepository() OrderRepository {
	return &orderRepository{db: database.DB}
}

func NewOrderItemRepository() OrderItemRepository {
	return &orderItemRepository{db: database.DB}
}

func NewOrderStatusRepository() OrderStatusRepository {
	return &orderStatusRepository{db: database.DB}
}

func NewOrderHistoryRepository() OrderHistoryRepository {
	return &orderHistoryRepository{db: database.DB}
}

// Order Repository Implementation
func (r *orderRepository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error
	return order, err
}

func (r *orderRepository) GetOrderByID(id uint) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("OrderItems").Preload("OrderHistories").First(&order, id).Error
	return order, err
}

func (r *orderRepository) UpdateOrder(order models.Order) (models.Order, error) {
	err := r.db.Save(&order).Error
	return order, err
}

func (r *orderRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&models.Order{}, id).Error
}

func (r *orderRepository) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Preload("OrderItems").Preload("OrderHistories").Find(&orders).Error
	return orders, err
}

// OrderItem Repository Implementation
func (r *orderItemRepository) CreateOrderItem(orderItem models.OrderItem) (models.OrderItem, error) {
	err := r.db.Create(&orderItem).Error
	return orderItem, err
}

func (r *orderItemRepository) GetOrderItemByID(id uint) (models.OrderItem, error) {
	var orderItem models.OrderItem
	err := r.db.First(&orderItem, id).Error
	return orderItem, err
}

func (r *orderItemRepository) UpdateOrderItem(orderItem models.OrderItem) (models.OrderItem, error) {
	err := r.db.Save(&orderItem).Error
	return orderItem, err
}

func (r *orderItemRepository) DeleteOrderItem(id uint) error {
	return r.db.Delete(&models.OrderItem{}, id).Error
}

// OrderStatus Repository Implementation
func (r *orderStatusRepository) CreateOrderStatus(orderStatus models.OrderStatus) (models.OrderStatus, error) {
	err := r.db.Create(&orderStatus).Error
	return orderStatus, err
}

func (r *orderStatusRepository) GetOrderStatusByID(id uint) (models.OrderStatus, error) {
	var orderStatus models.OrderStatus
	err := r.db.First(&orderStatus, id).Error
	return orderStatus, err
}

func (r *orderStatusRepository) UpdateOrderStatus(orderStatus models.OrderStatus) (models.OrderStatus, error) {
	err := r.db.Save(&orderStatus).Error
	return orderStatus, err
}

func (r *orderStatusRepository) DeleteOrderStatus(id uint) error {
	return r.db.Delete(&models.OrderStatus{}, id).Error
}

func (r *orderStatusRepository) GetOrderStatusByName(name string) (models.OrderStatus, error) {
	var orderStatus models.OrderStatus
	err := r.db.Where("status_name = ?", name).First(&orderStatus).Error
	return orderStatus, err
}

// OrderHistory Repository Implementation
func (r *orderHistoryRepository) CreateOrderHistory(orderHistory models.OrderHistory) (models.OrderHistory, error) {
	err := r.db.Create(&orderHistory).Error
	return orderHistory, err
}

func (r *orderHistoryRepository) GetOrderHistoryByID(id uint) (models.OrderHistory, error) {
	var orderHistory models.OrderHistory
	err := r.db.First(&orderHistory, id).Error
	return orderHistory, err
}

func (r *orderHistoryRepository) UpdateOrderHistory(orderHistory models.OrderHistory) (models.OrderHistory, error) {
	err := r.db.Save(&orderHistory).Error
	return orderHistory, err
}

func (r *orderHistoryRepository) DeleteOrderHistory(id uint) error {
	return r.db.Delete(&models.OrderHistory{}, id).Error
}

func (r *orderHistoryRepository) GetOrderHistoriesByOrderID(orderID uint) ([]models.OrderHistory, error) {
	var histories []models.OrderHistory
	err := r.db.Where("order_id = ?", orderID).Find(&histories).Error
	return histories, err
}
