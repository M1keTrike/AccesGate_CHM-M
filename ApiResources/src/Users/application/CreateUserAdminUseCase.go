package application

import (
	"api_resources/src/Users/domain"
	"api_resources/src/Users/domain/entities"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserAdminUseCase struct {
	Repo domain.UserRepository
}

func NewCreateUserAdminUseCase(repo domain.UserRepository) *CreateUserAdminUseCase {
	return &CreateUserAdminUseCase{Repo: repo}
}

func (uc *CreateUserAdminUseCase) Execute(user *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return uc.Repo.CreateUserAdmin(user)
}