package route

import (
	"product_management/api/handler"

	"github.com/labstack/echo/v4"
)

func NewProductRoute(e *echo.Echo, handler *handler.ProductHandler, middleware ...echo.MiddlewareFunc) {
	r := e.Group("/product")
	r.Use(middleware...)
	r.GET("", handler.ListProduct)
	r.POST("", handler.CreateProduct)
	r.GET("/:id", handler.GetProduct)
}
