package application

import (
    "api_resources/src/Users/domain"
	"api_resources/src/Users/domain/entities"
)

type UpdateUser struct {
    repo domain.UserRepository
}

func NewUpdateUser(repo domain.UserRepository) *UpdateUser {
    return &UpdateUser{repo: repo}
}

func (uc *UpdateUser) Execute(user *entities.User) error {
    return uc.repo.UpdateUser(user)
}