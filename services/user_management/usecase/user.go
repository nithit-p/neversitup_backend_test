package usecase

import (
	"user_management/domain"
	"user_management/repository"
)

type UserUsecase interface {
	ListUser() ([]domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
	GetUserById(id int) (*domain.User, error)
	CreateUser(username, email, firstName, lastName string) error
}

// Verify Interface
var _ UserUsecase = (*UserUsecaseImpl)(nil)

type UserUsecaseImpl struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecaseImpl {
	return &UserUsecaseImpl{
		repo: repo,
	}
}

func (uc *UserUsecaseImpl) ListUser() ([]domain.User, error) {
	users, err := uc.repo.ListUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUsecaseImpl) GetUserById(id int) (*domain.User, error) {
	user, err := uc.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uc *UserUsecaseImpl) GetUserByUsername(username string) (*domain.User, error) {
	user, err := uc.repo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uc *UserUsecaseImpl) CreateUser(username, email, firstName, lastName string) error {
	if err := uc.repo.CreateUser(username, email, firstName, lastName); err != nil {
		return err
	}
	return nil
}
