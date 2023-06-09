package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"order_management/api/handler"
	"order_management/api/route"
	"order_management/domain"
	"order_management/repository/postgresql"
	"order_management/service"
	"order_management/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	config := getConfig()

	// Establish a connection to database
	postgresDataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	postgresDB, err := sql.Open("postgres", postgresDataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return
	}
	defer postgresDB.Close()

	// Create repository
	orderRepo := postgresql.NewOrderRepository(postgresDB, config.ORDER_TABLE_NAME, config.ORDER_ITEM_TABLE_NAME)

	// Create httpservice
	productManagementService := service.NewProductManagementService(config.PRODUCT_MANAGEMENT_SERVICE_ADDRESS)

	// Create usecase
	orderUsecase := usecase.NewOrderUsecase(orderRepo, productManagementService)

	// Create handler
	orderHandler := handler.NewOrderHandler(orderUsecase)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JWTCustomClaims)
		},
		SigningKey: []byte(config.JWT_SECRET),
	}
	jwtMiddleware := echojwt.WithConfig(jwtConfig)

	// Create route
	route.NewOrderRoute(e, orderHandler, jwtMiddleware)
	route.NewOrderInternalRoute(e, orderHandler)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}

func getConfig() *config {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load the env vars in .env file")
	}

	return &config{
		PRODUCT_MANAGEMENT_SERVICE_ADDRESS: os.Getenv("PRODUCT_MANAGEMENT_SERVICE_ADDRESS"),

		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),

		ORDER_TABLE_NAME:      os.Getenv("ORDER_TABLE_NAME"),
		ORDER_ITEM_TABLE_NAME: os.Getenv("ORDER_ITEM_TABLE_NAME"),

		JWT_SECRET: os.Getenv("JWT_SECRET"),
		PORT:       os.Getenv("PORT"),
	}
}

type config struct {
	PRODUCT_MANAGEMENT_SERVICE_ADDRESS string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	ORDER_TABLE_NAME      string
	ORDER_ITEM_TABLE_NAME string

	JWT_SECRET string

	PORT string
}
