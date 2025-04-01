package application

import (
	"api_resources/src/AccessEvents/domain"
	"api_resources/src/AccessEvents/domain/entities"
)

type GetAccessEventsByDevice struct {
	Repo domain.AccessEventRepository
}

func NewGetAccessEventsByDevice(repo domain.AccessEventRepository) *GetAccessEventsByDevice {
	return &GetAccessEventsByDevice{Repo: repo}
}

func (g *GetAccessEventsByDevice) Execute(deviceID int) ([]entities.AccessEvent, error) {
	return g.Repo.GetByDevice(deviceID)
}
