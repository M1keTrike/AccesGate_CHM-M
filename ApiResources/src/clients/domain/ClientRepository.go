package domain

import (
	"api_resources/src/clients/domain/entities"
)

type ClientRepository interface {
	CreateClient(client *entities.Client) error
	GetAllClients() ([]entities.Client, error)
}
