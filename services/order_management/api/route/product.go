package route

import (
	"order_management/api/handler"

	"github.com/labstack/echo/v4"
)

func NewOrderRoute(e *echo.Echo, handler *handler.OrderHandler, middleware ...echo.MiddlewareFunc) {
	r := e.Group("/order")
	r.Use(middleware...)
	r.POST("", handler.CreateOrder)
	r.GET("/:id", handler.GetOrder)
	r.POST("/:id/cancel", handler.CancelOrder)
}

func NewOrderInternalRoute(e *echo.Echo, handler *handler.OrderHandler, middleware ...echo.MiddlewareFunc) {
	r := e.Group("/internal/order")
	r.Use(middleware...)
	r.GET("", handler.ListOrder)
	r.GET("/:id", handler.GetOrder)
	r.GET("/user/:id", handler.ListOrderByUserId)
}
