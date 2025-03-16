package application

import (
    "api_resources/src/Users/domain"
)

type DeleteUser struct {
    repo domain.UserRepository
}

func NewDeleteUser(repo domain.UserRepository) *DeleteUser {
    return &DeleteUser{repo: repo}
}

func (uc *DeleteUser) Execute(id int) error {
    return uc.repo.DeleteUser(id)
}