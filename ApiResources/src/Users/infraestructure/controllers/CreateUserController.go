package controllers

import (
	"api_resources/src/Users/application"
	"api_resources/src/Users/domain/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateUserRequest struct {
	Name          string `json:"name" example:"Admin"`
	Email         string `json:"email" example:"admin@example.com"`
	PasswordHash  string `json:"password_hash" example:"123456"`
	Role          string `json:"role" example:"admin"`
	FingerprintID int16  `json:"fingerprint_id" example:"0"`
	BiometricAuth bool   `json:"biometric_auth" example:"false"`
	CreatedBy     int    `json:"created_by" example:"1"`
}

type CreateUserController struct {
	useCase application.CreateUserUseCase
}

func NewCreateUserController(useCase application.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{useCase: useCase}
}

// CreateUser godoc
// @Summary Crea un nuevo usuario
// @Description Registra un usuario con nombre, email, contraseña y rol
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "Nuevo usuario"
// @Success 201 {object} entities.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func (c *CreateUserController) Execute(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	user := entities.User{
		Name:          req.Name,
		Email:         req.Email,
		PasswordHash:  req.PasswordHash,
		Role:          req.Role,
		FingerprintID: req.FingerprintID,
		BiometricAuth: req.BiometricAuth,
		CreatedBy:     req.CreatedBy,
	}

	err := c.useCase.Execute(user)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			ctx.JSON(http.StatusConflict, gin.H{"error": "El correo electrónico ya está registrado"})
			return
		}

		log.Printf("[CreateUser] Error inesperado: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
