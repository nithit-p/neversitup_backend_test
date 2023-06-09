package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"authentication/api/handler"
	"authentication/api/route"
	"authentication/repository/postgresql"
	"authentication/service"
	"authentication/usecase"

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
	authRepo := postgresql.NewAuthRepository(postgresDB, config.AUTH_TABLE_NAME)

	// Create httpservice
	userManagementService := service.NewUserManagementService(config.USER_MANAGEMENT_SERVICE_ADDRESS)

	// Create usecase
	authUsecase := usecase.NewAuthUsecase(authRepo, userManagementService)

	// Create handler
	authHandler := handler.NewAuthHandler(authUsecase, config.JWT_SECRET, config.JWT_EXPIRATION_TIME)

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

	// Create route
	route.NewAuthRoute(e, authHandler)

	e.Logger.Fatal(e.Start(":" + config.PORT))
}

func getConfig() *config {
	if err := godotenv.Load(); err != nil {
		log.Println("Failed to load the env vars in .env file")
	}
	jwtExpTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))
	if err != nil {
		log.Fatal("JWT_EXPIRATION_TIME invalid format")
	}
	return &config{

		USER_MANAGEMENT_SERVICE_ADDRESS: os.Getenv("USER_MANAGEMENT_SERVICE_ADDRESS"),

		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),

		AUTH_TABLE_NAME: os.Getenv("AUTH_TABLE_NAME"),

		JWT_SECRET:          os.Getenv("JWT_SECRET"),
		JWT_EXPIRATION_TIME: time.Duration(jwtExpTime) * time.Second,
		PORT:                os.Getenv("PORT"),
	}
}

type config struct {
	USER_MANAGEMENT_SERVICE_ADDRESS string

	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	AUTH_TABLE_NAME string

	JWT_SECRET          string
	JWT_EXPIRATION_TIME time.Duration

	PORT string
}
