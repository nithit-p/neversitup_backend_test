package usecase

import (
	"product_management/domain"
	"product_management/repository"
)

type ProductUsecase interface {
	GetAllProduct() ([]domain.Product, error)
	GetProductById(userId int) (*domain.Product, error)
	CreateProduct(name string, description string, price int) error
}

// Verify Interface
var _ ProductUsecase = (*ProductUsecaseImpl)(nil)

type ProductUsecaseImpl struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecaseImpl {
	return &ProductUsecaseImpl{
		repo: repo,
	}
}

func (uc *ProductUsecaseImpl) GetAllProduct() ([]domain.Product, error) {
	orders, err := uc.repo.ListProduct()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (uc *ProductUsecaseImpl) GetProductById(userId int) (*domain.Product, error) {
	product, err := uc.repo.GetProductById(userId)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (uc *ProductUsecaseImpl) CreateProduct(name string, description string, price int) error {
	if err := uc.repo.CreateProduct(name, description, price); err != nil {
		return err
	}
	return nil
}
