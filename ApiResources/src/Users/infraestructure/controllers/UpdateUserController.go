package controllers

import (
	"api_resources/src/Users/application"
	"api_resources/src/Users/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	useCase application.UpdateUser
}

func NewUpdateUserController(useCase application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{useCase: useCase}
}

// UpdateUser godoc
// @Summary Actualiza un usuario existente
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID del usuario"
// @Param user body entities.User true "Usuario actualizado"
// @Success 200 {object} entities.User
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [put]
func (c *UpdateUserController) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	user.ID = id

	if err := c.useCase.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el usuario"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}
