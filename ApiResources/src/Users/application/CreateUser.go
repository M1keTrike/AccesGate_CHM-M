package application

import (
	"api_resources/src/Users/domain"
	"api_resources/src/Users/domain/entities"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	Repo domain.UserRepository
}

func NewCreateUserUseCase(repo domain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{Repo: repo}
}

func (uc *CreateUserUseCase) Execute(user entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return uc.Repo.CreateUser(&user)
}
