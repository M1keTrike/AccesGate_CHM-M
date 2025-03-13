package domain

import "api_resources/src/Events/domain/entities"

type EventRepository interface {
	CreateEvent(event entities.Event) error
	GetEventByID(id int) (entities.Event, error)
	GetAllEvents() ([]entities.Event, error)
	UpdateEvent(event entities.Event) error
	DeleteEvent(id int) error
}
