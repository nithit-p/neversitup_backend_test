package handler

import (
	"log"
	"net/http"
	"product_management/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{
		usecase: usecase,
	}
}

func (handler *ProductHandler) ListProduct(c echo.Context) error {
	products, err := handler.usecase.GetAllProduct()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

func (handler *ProductHandler) GetProduct(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, "not found")
	}
	product, err := handler.usecase.GetProductById(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, product)
}

func (handler *ProductHandler) CreateProduct(c echo.Context) error {
	type productRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Price       int    `json:"price"`
	}
	newRequest := new(productRequest)
	if err := c.Bind(newRequest); err != nil {
		return err
	}

	if err := handler.usecase.CreateProduct(newRequest.Name, newRequest.Description, newRequest.Price); err != nil {
		return err
	}
	return c.String(http.StatusOK, "ok")
}
