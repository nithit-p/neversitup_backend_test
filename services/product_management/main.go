package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"product_management/api/handler"
	"product_management/api/route"
	"product_management/repository/postgresql"
	"product_management/usecase"

	"github.com/joho/godotenv"
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
	productRepo := postgresql.NewProductRepository(postgresDB, config.PRODUCT_TABLE_NAME)

	// Create usecase
	productUsecase := usecase.NewProductUsecase(productRepo)

	// Create handler
	productHandler := handler.NewProductHandler(productUsecase)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// jwtConfig := echojwt.Config{
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(domain.JWTCustomClaims)
	// 	},
	// 	SigningKey: []byte(config.JWT_SECRET),
	// }
	// jwtMiddleware := echojwt.WithConfig(jwtConfig)

	// Create route
	route.NewProductRoute(e, productHandler)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}

func getConfig() *config {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load the env vars in .env file")
	}

	return &config{
		PRODUCT_TABLE_NAME: os.Getenv("PRODUCT_TABLE_NAME"),

		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),

		JWT_SECRET: os.Getenv("JWT_SECRET"),
		PORT:       os.Getenv("PORT"),
	}
}

type config struct {
	PRODUCT_TABLE_NAME string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	JWT_SECRET string

	PORT string
}
