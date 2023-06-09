package domain

import "time"

type Product struct {
	ProductId   int       `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}
