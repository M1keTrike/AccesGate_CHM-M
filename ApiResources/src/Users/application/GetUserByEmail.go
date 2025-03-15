package application

import (
	"api_resources/src/Users/domain"
	"api_resources/src/Users/domain/entities"
)

type GetUserByEmail struct {
	repo domain.UserRepository
}

func NewGetUserByEmail(repo domain.UserRepository) *GetUserByEmail {
	return &GetUserByEmail{repo: repo}
}

func (uc *GetUserByEmail) Execute(email string) (*entities.User, error) {
	return uc.repo.GetUserByEmail(email)
}
