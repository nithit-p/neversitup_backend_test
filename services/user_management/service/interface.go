package service

import "user_management/domain"

type OrderManagementService interface {
	GetAllOrderByUserId(userId int) ([]domain.Order, error)
}
