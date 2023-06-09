package usecase

import (
	"authentication/domain"
	"authentication/repository"
	"authentication/service"
)

type AuthUsecase interface {
	Login(username string, password string) error
	GetUserAuthByUsername(username string) (*domain.Auth, error)
	CreateUser(username, email, password, firstName, lastName string) error
}

// Verify Interface
var _ AuthUsecase = (*AuthUsecaseImpl)(nil)

type AuthUsecaseImpl struct {
	repo                  repository.AuthRepository
	userManagementService service.UserManagementService
}

func NewAuthUsecase(repo repository.AuthRepository, userManagementService service.UserManagementService) *AuthUsecaseImpl {
	return &AuthUsecaseImpl{
		repo:                  repo,
		userManagementService: userManagementService,
	}
}

func (uc *AuthUsecaseImpl) Login(username string, password string) error {
	if err := uc.repo.Login(username, password); err != nil {
		return err
	}
	return nil
}

func (uc *AuthUsecaseImpl) GetUserAuthByUsername(username string) (*domain.Auth, error) {
	auth, err := uc.repo.GetAuthByUsername(username)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func (uc *AuthUsecaseImpl) CreateUser(username, email, password, firstName, lastName string) error {
	if err := uc.userManagementService.CreateUser(username, email, firstName, lastName); err != nil {
		return err
	}
	user, err := uc.userManagementService.GetUserByUsername(username)
	if err != nil {
		return err
	}
	if err := uc.repo.CreateAuth(user.ID, username, email, password); err != nil {
		return err
	}
	return nil
}
