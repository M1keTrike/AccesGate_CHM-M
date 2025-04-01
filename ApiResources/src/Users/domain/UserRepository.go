package domain

import (
	"api_resources/src/Users/domain/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id int) (entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetUsersByRole(role string) ([]entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id int) error
}
