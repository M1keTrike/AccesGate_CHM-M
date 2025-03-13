package application

import (
	"api_resources/src/Users/domain"
	"api_resources/src/Users/domain/entities"
)

type GetAllUsers struct {
	repo domain.UserRepository
}

func NewGetAllUsers(repo domain.UserRepository) *GetAllUsers {
	return &GetAllUsers{repo: repo}
}

func (uc *GetAllUsers) Execute() ([]entities.User, error) {
	return uc.repo.GetAllUsers()
}
