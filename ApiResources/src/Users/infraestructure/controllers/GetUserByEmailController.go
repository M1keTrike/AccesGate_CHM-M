package controllers

import (
	"api_resources/src/Users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)




type GetUserByEmailController struct {
	useCase application.GetUserByEmail
}

func NewGetUserByEmailController(useCase application.GetUserByEmail) *GetUserByEmailController {
	return &GetUserByEmailController{useCase: useCase}
}

// GetUserByEmail godoc
// @Summary Obtiene un usuario por email
// @Tags Users
// @Produce json
// @Param email query string true "Email del usuario"
// @Success 200 {object} entities.User
// @Failure 404 {object} map[string]string
// @Router /users/email [get]
func (c *GetUserByEmailController) Execute(ctx *gin.Context) {
	email := ctx.Query("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	user, err := c.useCase.Execute(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
