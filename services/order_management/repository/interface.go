package repository

import "order_management/domain"

type OrderRepository interface {
	ListOrder() ([]domain.Order, error)
	GetOrderById(id int) (*domain.Order, error)
	CreateOrder(userId int, items []domain.OrderItem) error
	ListOrderHistoryByUserId(userId int) ([]domain.Order, error)

	UpdateOrderStatus(id int, status string) error
}
