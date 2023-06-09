package route

import (
	"authentication/api/handler"

	"github.com/labstack/echo/v4"
)

func NewAuthRoute(e *echo.Echo, handler *handler.AuthHandler) {
	e.POST("/login", handler.Login)
	e.POST("/register", handler.Register)
}
