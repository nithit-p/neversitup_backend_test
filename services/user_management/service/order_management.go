package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"user_management/domain"
)

type OrderManagementServiceImpl struct {
	address string
}

func NewOrderManagementService(addr string) *OrderManagementServiceImpl {
	return &OrderManagementServiceImpl{
		address: addr,
	}
}

func (api *OrderManagementServiceImpl) GetAllOrderByUserId(userId int) ([]domain.Order, error) {
	baseURL, err := url.Parse(fmt.Sprintf("%s/internal/order/user/%d", api.address, userId))
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var orders []domain.Order
		json.NewDecoder(resp.Body).Decode(&orders)
		return orders, nil
	} else {
		return nil, fmt.Errorf("order-management-service error status:%d", resp.StatusCode)
	}
}
