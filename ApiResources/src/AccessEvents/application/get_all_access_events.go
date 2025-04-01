package application

import (
	"api_resources/src/AccessEvents/domain"
	"api_resources/src/AccessEvents/domain/entities"
)

type GetAllAccessEvents struct {
	Repo domain.AccessEventRepository
}

func NewGetAllAccessEvents(repo domain.AccessEventRepository) *GetAllAccessEvents {
	return &GetAllAccessEvents{Repo: repo}
}

func (g *GetAllAccessEvents) Execute() ([]entities.AccessEvent, error) {
	return g.Repo.GetAll()
}
