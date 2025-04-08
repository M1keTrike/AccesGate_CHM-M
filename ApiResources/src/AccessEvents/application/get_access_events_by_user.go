package application

import (
	"api_resources/src/AccessEvents/domain"
	"api_resources/src/AccessEvents/domain/entities"
)

type GetAccessEventsByUser struct {
	Repo domain.AccessEventRepository
}

func NewGetAccessEventsByUser(repo domain.AccessEventRepository) *GetAccessEventsByUser {
	return &GetAccessEventsByUser{Repo: repo}
}

func (g *GetAccessEventsByUser) Execute(userID int) ([]entities.AccessEvent, error) {
	return g.Repo.GetByUser(userID)
}
