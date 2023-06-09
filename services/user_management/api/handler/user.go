package handler

import (
	"net/http"
	"user_management/domain"
	"user_management/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase             usecase.UserUsecase
	usecaseOrderHistory usecase.OrderHistoryUsecase
}

func NewUserHandler(usecase usecase.UserUsecase, usecaseOrderHistory usecase.OrderHistoryUsecase) *UserHandler {
	return &UserHandler{
		usecase:             usecase,
		usecaseOrderHistory: usecaseOrderHistory,
	}
}

func (handler *UserHandler) ListUser(c echo.Context) error {
	users, err := handler.usecase.ListUser()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (handler *UserHandler) GetUserFormJWT(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*domain.JWTCustomClaims)
	name := claims.Name

	user, err := handler.usecase.GetUserByUsername(name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) GetUserByUsername(c echo.Context) error {
	username := c.Param("username")
	user, err := handler.usecase.GetUserByUsername(username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	newUser := new(domain.User)
	if err := c.Bind(newUser); err != nil {
		return err
	}
	if err := handler.usecase.CreateUser(newUser.Username, newUser.Email, newUser.FirstName, newUser.LastName); err != nil {
		return err
	}
	return c.String(http.StatusOK, "ok")
}

func (handler *UserHandler) ListOrderHistory(c echo.Context) error {
	userToken := c.Get("user").(*jwt.Token)
	claims := userToken.Claims.(*domain.JWTCustomClaims)

	orders, err := handler.usecaseOrderHistory.GetAllOrderHistoryByUserId(claims.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, orders)
}
