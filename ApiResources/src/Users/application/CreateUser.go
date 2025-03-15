package application

import (
    "api_resources/src/Users/domain"
	"api_resources/src/Users/domain/entities"
)

type CreateUser struct {
	repo domain.UserRepository
}

func NewCreateUser(repo domain.UserRepository) *CreateUser {
	return &CreateUser{repo: repo}
}

func (uc *CreateUser) Execute(user *entities.User) error {
	err := uc.repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}