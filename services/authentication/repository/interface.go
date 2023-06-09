package repository

import "authentication/domain"

type UserRepository interface {
	Login(username string, password string) error
	GetUserByUsername(username string) (domain.User, error)
	GetUserByID(id int) (domain.User, error)
	CreateUser(username, email, password string) error
}

type AuthRepository interface {
	Login(username string, password string) error
	GetAuthByUsername(username string) (domain.Auth, error)
	GetAuthByID(id int) (domain.Auth, error)
	CreateAuth(userId int, username, email, password string) error
}
