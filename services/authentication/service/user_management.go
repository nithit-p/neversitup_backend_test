package service

import (
	"authentication/domain"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type UserManagementServiceImpl struct {
	address string
}

func NewUserManagementService(addr string) *UserManagementServiceImpl {
	return &UserManagementServiceImpl{
		address: addr,
	}
}

func (api *UserManagementServiceImpl) GetUserByUsername(username string) (*domain.User, error) {
	baseURL, err := url.Parse(fmt.Sprintf("%s/internal/user/%s", api.address, username))
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var user domain.User
		json.NewDecoder(resp.Body).Decode(&user)
		return &user, nil
	} else {
		return nil, fmt.Errorf("user-management-service error status:%d", resp.StatusCode)
	}
}

func (api *UserManagementServiceImpl) CreateUser(username, email, firstName, lastName string) error {
	type userRequest struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	requestData := userRequest{
		Username:  username,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
	requestDataJson, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	baseURL, err := url.Parse(fmt.Sprintf("%s/internal/user", api.address))
	if err != nil {
		return err
	}
	resp, err := http.Post(baseURL.String(), "application/json", bytes.NewBuffer(requestDataJson))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("user-management-service error status:%d", resp.StatusCode)
	}
}
