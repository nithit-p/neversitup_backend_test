package domain

import "time"

type Order struct {
	OrderId     int         `json:"order_id"`
	UserId      int         `json:"user_id"`
	TotalAmount int         `json:"total_amount"`
	OrderDate   time.Time   `json:"order_date"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	OrderItemId        int    `json:"order_item_id"`
	OrderId            int    `json:"order_id"`
	ProductId          int    `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductPrice       int    `json:"product_price"`
	Quantity           int    `json:"quantity"`
	Price              int    `json:"price"`
}

type Product struct {
	ProductId   int       `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}
