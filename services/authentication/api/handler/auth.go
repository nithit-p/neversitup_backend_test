package handler

import (
	"log"
	"net/http"
	"time"

	"authentication/domain"
	"authentication/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	usecase    usecase.AuthUsecase
	jwtSecret  []byte
	jwtExpTime time.Duration
}

func NewAuthHandler(usecase usecase.AuthUsecase, jwtSecret string, jwtExpTime time.Duration) *AuthHandler {
	return &AuthHandler{
		usecase:    usecase,
		jwtSecret:  []byte(jwtSecret),
		jwtExpTime: jwtExpTime,
	}
}

func (handler *AuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if err := handler.usecase.Login(username, password); err != nil {
		log.Println(err)
		return echo.ErrUnauthorized
	}

	// Get user data
	user, err := handler.usecase.GetUserAuthByUsername(username)
	if err != nil {
		log.Println(err)
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &domain.JWTCustomClaims{
		ID:    user.UserId,
		Name:  user.Username,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(handler.jwtExpTime)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(handler.jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (handler *AuthHandler) Register(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	rePassword := c.FormValue("re-password")
	email := c.FormValue("email")
	firstName := c.FormValue("first_name")
	lastName := c.FormValue("last_name")

	// Check password
	if password != rePassword {
		return echo.ErrBadRequest
	}

	if err := handler.usecase.CreateUser(username, email, password, firstName, lastName); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}

	// Get user data
	userAuth, err := handler.usecase.GetUserAuthByUsername(username)
	if err != nil {
		log.Println(err)
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &domain.JWTCustomClaims{
		ID:    userAuth.UserId,
		Name:  userAuth.Username,
		Email: userAuth.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(handler.jwtExpTime)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString(handler.jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
