package application

import (
	"api_resources/src/AccessEvents/domain"
	"api_resources/src/AccessEvents/domain/entities"
)

type GetAccessEventsByFront struct {
	Repo domain.AccessEventRepository
}

func NewGetAccessEventsByFront(repo domain.AccessEventRepository) *GetAccessEventsByFront {
	return &GetAccessEventsByFront{Repo: repo}
}

func (g *GetAccessEventsByFront) Execute(frontID int) ([]entities.AccessEvent, error) {
	return g.Repo.GetByFront(frontID)
}
