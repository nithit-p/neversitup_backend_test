package domain

import "time"

type Order struct {
	OrderId     int         `json:"order_id"`
	UserId      int         `json:"user_id"`
	TotalAmount int         `json:"total_amount"`
	Status      string      `json:"status"`
	OrderDate   time.Time   `json:"order_date"`
	Items       []OrderItem `json:"items"`
}

type OrderItem struct {
	OrderItemId int `json:"order_item_id"`
	OrderId     int `json:"order_id"`
	ProductId   int `json:"product_id"`
	Quantity    int `json:"quantity"`
	Price       int `json:"price"`
}
