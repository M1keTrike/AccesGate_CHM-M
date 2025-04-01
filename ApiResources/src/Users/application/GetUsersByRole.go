package application

import (
    "api_resources/src/Users/domain"
    "api_resources/src/Users/domain/entities"
)

type GetUsersByRole struct {
    repo domain.UserRepository
}

func NewGetUsersByRole(repo domain.UserRepository) *GetUsersByRole {
    return &GetUsersByRole{repo: repo}
}

func (uc *GetUsersByRole) Execute(role string) ([]entities.User, error) {
    return uc.repo.GetUsersByRole(role)
}