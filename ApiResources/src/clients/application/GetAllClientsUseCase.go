package application

import (
	"api_resources/src/clients/domain"
	"api_resources/src/clients/domain/entities"
)

type GetAllClientsUseCase struct {
	repo domain.ClientRepository
}

func NewGetAllClientsUseCase(repo domain.ClientRepository) *GetAllClientsUseCase {
	return &GetAllClientsUseCase{repo: repo}
}

func (uc *GetAllClientsUseCase) Execute() ([]entities.Client, error) {
	return uc.repo.GetAllClients()
}
