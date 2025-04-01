package domain

import "api_resources/src/AccessEvents/domain/entities"

type AccessEventRepository interface {
	Create(event *entities.AccessEvent) error
	GetAll() ([]entities.AccessEvent, error)
	GetByUser(userID int) ([]entities.AccessEvent, error)
	GetByDevice(deviceID int) ([]entities.AccessEvent, error)
	GetByFront(frontID int) ([]entities.AccessEvent, error)
	GetByID(id int) (*entities.AccessEvent, error)
	Update(event *entities.AccessEvent) error
	Delete(id int) error
}
