package route

import (
	"user_management/api/handler"

	"github.com/labstack/echo/v4"
)

func NewUserRoute(e *echo.Echo, handler *handler.UserHandler, middleware ...echo.MiddlewareFunc) {
	r := e.Group("/user")
	r.Use(middleware...)
	r.GET("", handler.GetUserFormJWT)
	r.GET("/order", handler.ListOrderHistory)
}

func NewUserInternalRoute(e *echo.Echo, handler *handler.UserHandler, middleware ...echo.MiddlewareFunc) {
	r := e.Group("/internal/user")
	r.Use(middleware...)
	r.GET("", handler.ListUser)
	r.POST("", handler.CreateUser)
	r.GET("/:username", handler.GetUserByUsername)
}
