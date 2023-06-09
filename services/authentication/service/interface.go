package service

import "authentication/domain"

type UserManagementService interface {
	GetUserByUsername(username string) (*domain.User, error)
	CreateUser(username, email, firstName, lastName string) error
}
