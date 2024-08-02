package services

import (
	"github.com/rmarmolejo90/candles_api/internal/models"
	"github.com/rmarmolejo90/candles_api/internal/repository"
)

// OrderService defines the methods for interacting with Order data
type OrderService interface {
	CreateOrder(order models.Order) (models.Order, error)
	GetOrderByID(id uint) (models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(id uint) error
	GetOrdersByUserID(userID uint) ([]models.Order, error)
}

// OrderItemService defines the methods for interacting with OrderItem data
type OrderItemService interface {
	CreateOrderItem(orderItem models.OrderItem) (models.OrderItem, error)
	GetOrderItemByID(id uint) (models.OrderItem, error)
	UpdateOrderItem(orderItem models.OrderItem) (models.OrderItem, error)
	DeleteOrderItem(id uint) error
}

// OrderStatusService defines the methods for interacting with OrderStatus data
type OrderStatusService interface {
	CreateOrderStatus(orderStatus models.OrderStatus) (models.OrderStatus, error)
	GetOrderStatusByID(id uint) (models.OrderStatus, error)
	UpdateOrderStatus(orderStatus models.OrderStatus) (models.OrderStatus, error)
	DeleteOrderStatus(id uint) error
	GetOrderStatusByName(name string) (models.OrderStatus, error)
}

// OrderHistoryService defines the methods for interacting with OrderHistory data
type OrderHistoryService interface {
	CreateOrderHistory(orderHistory models.OrderHistory) (models.OrderHistory, error)
	GetOrderHistoryByID(id uint) (models.OrderHistory, error)
	UpdateOrderHistory(orderHistory models.OrderHistory) (models.OrderHistory, error)
	DeleteOrderHistory(id uint) error
	GetOrderHistoriesByOrderID(orderID uint) ([]models.OrderHistory, error)
}

// Order service implementations
type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{orderRepo: orderRepo}
}

func (s *orderService) CreateOrder(order models.Order) (models.Order, error) {
	// validation can be added here
	return s.orderRepo.CreateOrder(order)
}

func (s *orderService) GetOrderByID(id uint) (models.Order, error) {
	return s.orderRepo.GetOrderByID(id)
}

func (s *orderService) UpdateOrder(order models.Order) (models.Order, error) {
	return s.orderRepo.UpdateOrder(order)
}

func (s *orderService) DeleteOrder(id uint) error {
	return s.orderRepo.DeleteOrder(id)
}

func (s *orderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	return s.orderRepo.GetOrdersByUserID(userID)
}

// OrderItem service implementations
type orderItemService struct {
	orderItemRepo repository.OrderItemRepository
}

func NewOrderItemService(orderItemRepo repository.OrderItemRepository) OrderItemService {
	return &orderItemService{orderItemRepo: orderItemRepo}
}

func (s *orderItemService) CreateOrderItem(orderItem models.OrderItem) (models.OrderItem, error) {
	// validation can be added here
	return s.orderItemRepo.CreateOrderItem(orderItem)
}

func (s *orderItemService) GetOrderItemByID(id uint) (models.OrderItem, error) {
	return s.orderItemRepo.GetOrderItemByID(id)
}

func (s *orderItemService) UpdateOrderItem(orderItem models.OrderItem) (models.OrderItem, error) {
	return s.orderItemRepo.UpdateOrderItem(orderItem)
}

func (s *orderItemService) DeleteOrderItem(id uint) error {
	return s.orderItemRepo.DeleteOrderItem(id)
}

// OrderStatus service implementations
type orderStatusService struct {
	orderStatusRepo repository.OrderStatusRepository
}

func NewOrderStatusService(orderStatusRepo repository.OrderStatusRepository) OrderStatusService {
	return &orderStatusService{orderStatusRepo: orderStatusRepo}
}

func (s *orderStatusService) CreateOrderStatus(orderStatus models.OrderStatus) (models.OrderStatus, error) {
	return s.orderStatusRepo.CreateOrderStatus(orderStatus)
}

func (s *orderStatusService) GetOrderStatusByID(id uint) (models.OrderStatus, error) {
	return s.orderStatusRepo.GetOrderStatusByID(id)
}

func (s *orderStatusService) UpdateOrderStatus(orderStatus models.OrderStatus) (models.OrderStatus, error) {
	return s.orderStatusRepo.UpdateOrderStatus(orderStatus)
}

func (s *orderStatusService) DeleteOrderStatus(id uint) error {
	return s.orderStatusRepo.DeleteOrderStatus(id)
}

func (s *orderStatusService) GetOrderStatusByName(name string) (models.OrderStatus, error) {
	return s.orderStatusRepo.GetOrderStatusByName(name)
}

// OrderHistory service implementations
type orderHistoryService struct {
	orderHistoryRepo repository.OrderHistoryRepository
}

func NewOrderHistoryService(orderHistoryRepo repository.OrderHistoryRepository) OrderHistoryService {
	return &orderHistoryService{orderHistoryRepo: orderHistoryRepo}
}

func (s *orderHistoryService) CreateOrderHistory(orderHistory models.OrderHistory) (models.OrderHistory, error) {
	return s.orderHistoryRepo.CreateOrderHistory(orderHistory)
}

func (s *orderHistoryService) GetOrderHistoryByID(id uint) (models.OrderHistory, error) {
	return s.orderHistoryRepo.GetOrderHistoryByID(id)
}

func (s *orderHistoryService) UpdateOrderHistory(orderHistory models.OrderHistory) (models.OrderHistory, error) {
	return s.orderHistoryRepo.UpdateOrderHistory(orderHistory)
}

func (s *orderHistoryService) DeleteOrderHistory(id uint) error {
	return s.orderHistoryRepo.DeleteOrderHistory(id)
}

func (s *orderHistoryService) GetOrderHistoriesByOrderID(orderID uint) ([]models.OrderHistory, error) {
	return s.orderHistoryRepo.GetOrderHistoriesByOrderID(orderID)
}
