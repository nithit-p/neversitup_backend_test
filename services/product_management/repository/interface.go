package repository

import "product_management/domain"

type ProductRepository interface {
	ListProduct() ([]domain.Product, error)
	GetProductById(id int) (*domain.Product, error)
	CreateProduct(name string, description string, price int) error
}
