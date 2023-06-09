package postgresql

import (
	"database/sql"
	"fmt"

	"user_management/domain"
	"user_management/repository"
)

// Verify Interface
var _ repository.OrderHistoryRepository = (*OrderHistoryRepositoryImpl)(nil)

type OrderHistoryRepositoryImpl struct {
	db                 *sql.DB
	orderTableName     string
	orderItemTableName string
}

func NewOrderHistoryRepository(db *sql.DB, orderTableName string, orderItemTableName string) *OrderHistoryRepositoryImpl {
	return &OrderHistoryRepositoryImpl{
		db:                 db,
		orderTableName:     orderTableName,
		orderItemTableName: orderItemTableName,
	}
}

func (repo *OrderHistoryRepositoryImpl) ListOrderHistoryByUserId(userId int) ([]domain.Order, error) {
	var orders []domain.Order

	queryStr := fmt.Sprintf("SELECT order_id, user_id, total_amount, order_date FROM %s WHERE user_id=$1 ORDER BY order_date DESC", repo.orderTableName)
	rows, err := repo.db.Query(queryStr, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order domain.Order
		err := rows.Scan(&order.OrderId, &order.UserId, &order.TotalAmount, &order.OrderDate)
		if err != nil {
			return nil, err
		}
		orderItemsQueryStr := fmt.Sprintf("SELECT order_item_id, order_id, product_id, quantity, price FROM %s WHERE order_id=$1", repo.orderItemTableName)
		itemsRows, err := repo.db.Query(orderItemsQueryStr, order.OrderId)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		defer itemsRows.Close()

		for itemsRows.Next() {
			var item domain.OrderItem
			err := itemsRows.Scan(&item.OrderItemId, &item.OrderId, &item.ProductId, &item.Quantity, &item.Price)
			if err != nil {
				return nil, err
			}
			order.Items = append(order.Items, item)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
