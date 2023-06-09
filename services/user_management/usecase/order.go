package usecase

import (
	"user_management/domain"
	"user_management/service"
)

type OrderHistoryUsecase interface {
	GetAllOrderHistoryByUserId(userId int) ([]domain.Order, error)
}

// Verify Interface
var _ OrderHistoryUsecase = (*OrderHistoryUsecaseImpl)(nil)

type OrderHistoryUsecaseImpl struct {
	service service.OrderManagementService
}

func NewOrderHistoryUsecase(service service.OrderManagementService) *OrderHistoryUsecaseImpl {
	return &OrderHistoryUsecaseImpl{
		service: service,
	}
}

func (uc *OrderHistoryUsecaseImpl) GetAllOrderHistoryByUserId(userId int) ([]domain.Order, error) {
	orders, err := uc.service.GetAllOrderByUserId(userId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
