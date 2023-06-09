package service

import "order_management/domain"

type ProductManagementService interface {
	GetProductById(id int) (*domain.Product, error)
}
