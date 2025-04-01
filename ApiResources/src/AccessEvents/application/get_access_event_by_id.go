package application

import (
	"api_resources/src/AccessEvents/domain"
	"api_resources/src/AccessEvents/domain/entities"
)

type GetAccessEventByID struct {
	Repo domain.AccessEventRepository
}

func NewGetAccessEventByID(repo domain.AccessEventRepository) *GetAccessEventByID {
	return &GetAccessEventByID{Repo: repo}
}

func (g *GetAccessEventByID) Execute(id int) (*entities.AccessEvent, error) {
	return g.Repo.GetByID(id)
}
