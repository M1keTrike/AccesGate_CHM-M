package application

import (
	"api_resources/src/Users/domain"
	"api_resources/src/Users/domain/entities"
)

type GetUsersByCreatedByUseCase struct {
	userRepository domain.UserRepository
}

func NewGetUsersByCreatedByUseCase(userRepository domain.UserRepository) *GetUsersByCreatedByUseCase {
	return &GetUsersByCreatedByUseCase{
		userRepository: userRepository,
	}
}

func (uc *GetUsersByCreatedByUseCase) Execute(createdBy int) ([]entities.User, error) {
	return uc.userRepository.GetUsersByCreatedBy(createdBy)
}
