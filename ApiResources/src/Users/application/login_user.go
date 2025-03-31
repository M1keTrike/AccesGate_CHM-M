package application

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"api_resources/src/Users/domain"
)

var JwtSecret = []byte("your_secret_key") // cambia esto por una env var segura

type LoginUseCase struct {
	Repo domain.UserRepository
}

func NewLoginUseCase(repo domain.UserRepository) *LoginUseCase {
	return &LoginUseCase{Repo: repo}
}

func (uc *LoginUseCase) Execute(email, password string) (string, error) {
	user, err := uc.Repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("credenciales inválidas")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
