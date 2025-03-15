package application

import (
	"api_resources/src/clients/domain"
	"api_resources/src/clients/domain/entities"
)

type CreateClientUseCase struct {
	repo domain.ClientRepository
}

func NewCreateClientUseCase(repo domain.ClientRepository) *CreateClientUseCase {
	return &CreateClientUseCase{repo: repo}
}

func (uc *CreateClientUseCase) Execute(client *entities.Client) error {
	return uc.repo.CreateClient(client)
}
