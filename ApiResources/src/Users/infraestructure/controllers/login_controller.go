package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"api_resources/src/Users/application"
)

// LoginRequest representa el cuerpo de la solicitud de login
type LoginRequest struct {
	Email    string `json:"email" example:"admin@example.com"`
	Password string `json:"password" example:"123456"`
}

type LoginController struct {
	LoginUC *application.LoginUseCase
}

func NewLoginController(loginUC *application.LoginUseCase) *LoginController {
	return &LoginController{LoginUC: loginUC}
}

// LoginUser godoc
// @Summary Inicia sesión de usuario
// @Description Valida credenciales y retorna un token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Credenciales"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /users/login [post]
func (c *LoginController) Execute(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	token, err := c.LoginUC.Execute(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
