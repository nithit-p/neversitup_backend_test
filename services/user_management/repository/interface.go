package repository

import "user_management/domain"

type UserRepository interface {
	ListUser() ([]domain.User, error)
	GetUserByUsername(username string) (domain.User, error)
	GetUserByID(id int) (domain.User, error)
	CreateUser(username, email, firstName, lastName string) error
}

type OrderHistoryRepository interface {
	ListOrderHistoryByUserId(userId int) ([]domain.Order, error)
}
