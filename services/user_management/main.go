package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"user_management/api/handler"
	"user_management/api/route"
	"user_management/domain"
	"user_management/repository/postgresql"
	"user_management/service"
	"user_management/usecase"

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
	userRepo := postgresql.NewUserRepository(postgresDB, config.USER_TABLE_NAME)

	// Create httpservice
	orderManagementService := service.NewOrderManagementService(config.ORDER_MANAGEMENT_SERVICE_ADDRESS)

	// Create usecase
	userUsecase := usecase.NewUserUsecase(userRepo)
	orderHistoryUsecase := usecase.NewOrderHistoryUsecase(orderManagementService)

	// Create handler
	userHandler := handler.NewUserHandler(userUsecase, orderHistoryUsecase)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// JWT Middleware
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JWTCustomClaims)
		},
		SigningKey: []byte(config.JWT_SECRET),
	}
	jwtMiddleware := echojwt.WithConfig(jwtConfig)

	// ApiKey Middleware
	// apiKeyMiddleware := middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
	// 	KeyLookup: "query:api-key",
	// 	Validator: func(key string, c echo.Context) (bool, error) {
	// 		return key == config.ORDER_MANAGEMENT_API_KEY, nil
	// 	},
	// })

	// Create route
	route.NewUserRoute(e, userHandler, jwtMiddleware)
	// route.NewUserInternalRoute(e, userHandler, apiKeyMiddleware)
	route.NewUserInternalRoute(e, userHandler)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}

func getConfig() *config {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load the env vars in .env file")
	}

	return &config{
		ORDER_MANAGEMENT_SERVICE_ADDRESS: os.Getenv("ORDER_MANAGEMENT_SERVICE_ADDRESS"),

		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),

		USER_TABLE_NAME: os.Getenv("USER_TABLE_NAME"),

		JWT_SECRET: os.Getenv("JWT_SECRET"),
		PORT:       os.Getenv("PORT"),

		// ORDER_MANAGEMENT_API_KEY: os.Getenv("ORDER_MANAGEMENT_API_KEY"),
	}
}

type config struct {
	ORDER_MANAGEMENT_SERVICE_ADDRESS string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	USER_TABLE_NAME string

	JWT_SECRET string

	PORT string

	// ORDER_MANAGEMENT_API_KEY string
}
