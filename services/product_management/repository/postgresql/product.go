package postgresql

import (
	"database/sql"
	"fmt"

	"product_management/domain"
	"product_management/repository"
)

// Verify Interface
var _ repository.ProductRepository = (*ProductRepositoryImpl)(nil)

type ProductRepositoryImpl struct {
	db        *sql.DB
	tableName string
}

func NewProductRepository(db *sql.DB, tableName string) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (repo *ProductRepositoryImpl) ListProduct() ([]domain.Product, error) {
	products := make([]domain.Product, 0)

	queryStr := fmt.Sprintf("SELECT product_id, name, description, price, created_at FROM %s", repo.tableName)
	rows, err := repo.db.Query(queryStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.ProductId, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductRepositoryImpl) GetProductById(id int) (*domain.Product, error) {
	var product domain.Product
	queryStr := fmt.Sprintf("SELECT product_id, name, description, price, created_at FROM %s WHERE product_id=$1", repo.tableName)
	err := repo.db.QueryRow(queryStr, id).Scan(&product.ProductId, &product.Name, &product.Description, &product.Price, &product.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (repo *ProductRepositoryImpl) CreateProduct(name string, description string, price int) error {
	// Execute the SQL statement
	queryStr := fmt.Sprintf("INSERT INTO %s (name, description, price, created_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP)", repo.tableName)
	_, err := repo.db.Exec(queryStr, name, description, price)
	if err != nil {
		return err
	}

	return nil
}
