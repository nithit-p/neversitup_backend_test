package postgresql

import (
	"database/sql"
	"fmt"

	"order_management/domain"
	"order_management/repository"
)

// Verify Interface
var _ repository.OrderRepository = (*OrderRepositoryImpl)(nil)

type OrderRepositoryImpl struct {
	db            *sql.DB
	tableName     string
	tableItemName string
}

func NewOrderRepository(db *sql.DB, tableName string, tableItemName string) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{
		db:            db,
		tableName:     tableName,
		tableItemName: tableItemName,
	}
}

func (repo *OrderRepositoryImpl) ListOrder() ([]domain.Order, error) {
	orders := make([]domain.Order, 0)

	queryStr := fmt.Sprintf("SELECT order_id, user_id, total_amount, order_date, status FROM %s ORDER BY order_date DESC", repo.tableName)
	rows, err := repo.db.Query(queryStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order domain.Order
		err := rows.Scan(&order.OrderId, &order.UserId, &order.TotalAmount, &order.OrderDate, &order.Status)
		if err != nil {
			return nil, err
		}
		orderItemsQueryStr := fmt.Sprintf("SELECT order_item_id, order_id, product_id, quantity, price FROM %s WHERE order_id=$1", repo.tableItemName)
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

func (repo *OrderRepositoryImpl) GetOrderById(id int) (*domain.Order, error) {
	var order domain.Order

	queryStr := fmt.Sprintf("SELECT order_id, user_id, total_amount, order_date, status FROM %s WHERE order_id=$1", repo.tableName)
	if err := repo.db.QueryRow(queryStr, id).Scan(&order.OrderId, &order.UserId, &order.TotalAmount, &order.OrderDate, &order.Status); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	orderItemsQueryStr := fmt.Sprintf("SELECT order_item_id, order_id, product_id, quantity, price FROM %s WHERE order_id=$1", repo.tableItemName)
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

	return &order, nil
}

func (repo *OrderRepositoryImpl) CreateOrder(userId int, items []domain.OrderItem) error {
	totalAmount := 0
	for i := 0; i < len(items); i++ {
		totalAmount += items[i].Price * items[i].Quantity
	}
	// Execute the SQL statement
	queryStr := fmt.Sprintf("INSERT INTO %s (user_id, total_amount, status, order_date) VALUES ($1, $2, $3, CURRENT_TIMESTAMP)", repo.tableName)
	_, err := repo.db.Exec(queryStr, userId, totalAmount, "pending")
	if err != nil {
		return err
	}
	var orderId int
	orderItemsQueryStr := fmt.Sprintf("SELECT order_id FROM %s WHERE user_id=$1 AND total_amount=$2 AND status=$3 ORDER BY order_date DESC", repo.tableName)
	if err := repo.db.QueryRow(orderItemsQueryStr, userId, totalAmount, "pending").Scan(&orderId); err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}
	for i := 0; i < len(items); i++ {
		queryStr := fmt.Sprintf("INSERT INTO %s (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", repo.tableItemName)
		_, err := repo.db.Exec(queryStr, orderId, items[i].ProductId, items[i].Quantity, items[i].Price)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *OrderRepositoryImpl) ListOrderHistoryByUserId(userId int) ([]domain.Order, error) {
	orders := make([]domain.Order, 0)

	queryStr := fmt.Sprintf("SELECT order_id, user_id, total_amount, order_date, status FROM %s WHERE user_id=$1 ORDER BY order_date DESC", repo.tableName)
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
		err := rows.Scan(&order.OrderId, &order.UserId, &order.TotalAmount, &order.OrderDate, &order.Status)
		if err != nil {
			return nil, err
		}
		orderItemsQueryStr := fmt.Sprintf("SELECT order_item_id, order_id, product_id, quantity, price FROM %s WHERE order_id=$1", repo.tableItemName)
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

func (repo *OrderRepositoryImpl) UpdateOrderStatus(orderId int, status string) error {
	// Execute the SQL statement
	queryStr := fmt.Sprintf("UPDATE %s SET status = $1 WHERE order_id = $2", repo.tableName)
	_, err := repo.db.Exec(queryStr, status, orderId)
	if err != nil {
		return err
	}

	return nil
}
