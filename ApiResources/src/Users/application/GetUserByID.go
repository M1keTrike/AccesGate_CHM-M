package application

import (
	"api_resources/src/Users/domain"
	"api_resources/src/Users/domain/entities"
)

type GetUserByID struct {
	repo domain.UserRepository
}

func NewGetUserByID(repo domain.UserRepository) *GetUserByID {
	return &GetUserByID{repo: repo}
}

func (uc *GetUserByID) Execute(id int) (*entities.User, error) {
	User, err := uc.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}	
	return &User, nil
}
