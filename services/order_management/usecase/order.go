package usecase

import (
	"order_management/domain"
	"order_management/repository"
	"order_management/service"
)

type OrderUsecase interface {
	GetAllOrder() ([]domain.Order, error)
	GetOrderById(orderId int) (*domain.Order, error)
	CreateOrder(userId int, items []domain.OrderItem) error
	UpdateOrderStatus(orderId int, status string) error
	GetAllOrderHistoryByUserId(userId int) ([]domain.Order, error)
	GetProductByID(productId int) (*domain.Product, error)
}

// Verify Interface
var _ OrderUsecase = (*OrderUsecaseImpl)(nil)

type OrderUsecaseImpl struct {
	repo                     repository.OrderRepository
	productManagementService service.ProductManagementService
}

func NewOrderUsecase(repo repository.OrderRepository, productManagementService service.ProductManagementService) *OrderUsecaseImpl {
	return &OrderUsecaseImpl{
		repo:                     repo,
		productManagementService: productManagementService,
	}
}

func (uc *OrderUsecaseImpl) GetAllOrder() ([]domain.Order, error) {
	orders, err := uc.repo.ListOrder()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (uc *OrderUsecaseImpl) GetOrderById(orderId int) (*domain.Order, error) {
	order, err := uc.repo.GetOrderById(orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (uc *OrderUsecaseImpl) CreateOrder(userId int, items []domain.OrderItem) error {
	if err := uc.repo.CreateOrder(userId, items); err != nil {
		return err
	}
	return nil
}

func (uc *OrderUsecaseImpl) GetAllOrderHistoryByUserId(userId int) ([]domain.Order, error) {
	orders, err := uc.repo.ListOrderHistoryByUserId(userId)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (uc *OrderUsecaseImpl) UpdateOrderStatus(orderId int, status string) error {
	if err := uc.repo.UpdateOrderStatus(orderId, status); err != nil {
		return err
	}
	return nil
}

func (uc *OrderUsecaseImpl) GetProductByID(productId int) (*domain.Product, error) {
	product, err := uc.productManagementService.GetProductById(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
