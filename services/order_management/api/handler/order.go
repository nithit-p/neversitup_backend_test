package handler

import (
	"net/http"
	"order_management/domain"
	"order_management/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	usecase usecase.OrderUsecase
}

func NewOrderHandler(usecase usecase.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		usecase: usecase,
	}
}

func (handler *OrderHandler) ListOrder(c echo.Context) error {
	orders, err := handler.usecase.GetAllOrder()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orders)
}

func (handler *OrderHandler) ListOrderByUserId(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, "")
	}
	orders, err := handler.usecase.GetAllOrderHistoryByUserId(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orders)
}

func (handler *OrderHandler) GetOrder(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, "")
	}
	order, err := handler.usecase.GetOrderById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, order)
}

func (handler *OrderHandler) CreateOrder(c echo.Context) error {
	type orderItemRequest struct {
		ProductId int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}
	type orderRequest struct {
		Items []orderItemRequest `json:"items"`
	}

	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*domain.JWTCustomClaims)
	id := claims.ID

	newRequest := new(orderRequest)
	if err := c.Bind(newRequest); err != nil {
		return err
	}

	orderItems := make([]domain.OrderItem, 0)
	for i := 0; i < len(newRequest.Items); i++ {
		if newRequest.Items[i].Quantity == 0 {
			continue
		}
		product, err := handler.usecase.GetProductByID(newRequest.Items[i].ProductId)
		if err != nil {
			return err
		}
		orderItems = append(orderItems, domain.OrderItem{
			ProductId: product.ProductId,
			Price:     product.Price,
			Quantity:  newRequest.Items[i].Quantity,
		})
	}

	if err := handler.usecase.CreateOrder(id, orderItems); err != nil {
		return err
	}
	return c.String(http.StatusOK, "ok")
}

func (handler *OrderHandler) CancelOrder(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*domain.JWTCustomClaims)
	userId := claims.ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, "")
	}
	order, err := handler.usecase.GetOrderById(id)
	if err != nil {
		return err
	}
	if order.UserId != userId {
		return c.JSON(http.StatusForbidden, "ok")
	}
	if err := handler.usecase.UpdateOrderStatus(id, "canceled"); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "ok")
}
