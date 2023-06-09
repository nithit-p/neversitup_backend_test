package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"order_management/domain"
)

type ProductManagementServiceImpl struct {
	address string
}

func NewProductManagementService(addr string) *ProductManagementServiceImpl {
	return &ProductManagementServiceImpl{
		address: addr,
	}
}

func (api *ProductManagementServiceImpl) GetProductById(id int) (*domain.Product, error) {
	baseURL, err := url.Parse(fmt.Sprintf("%s/product/%d", api.address, id))
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var product domain.Product
		json.NewDecoder(resp.Body).Decode(&product)
		return &product, nil
	} else {
		return nil, fmt.Errorf("user-management-service error status:%d", resp.StatusCode)
	}
}
