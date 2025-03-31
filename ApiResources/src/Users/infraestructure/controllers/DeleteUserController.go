package controllers

import (
	"api_resources/src/Users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteUserController struct {
	useCase application.DeleteUser
}

func NewDeleteUserController(useCase application.DeleteUser) *DeleteUserController {
	return &DeleteUserController{useCase: useCase}
}

// DeleteUser godoc
// @Summary Elimina un usuario por ID
// @Description Borra un usuario existente utilizando su ID
// @Tags Users
// @Produce json
// @Param id path int true "ID del usuario"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [delete]
func (c *DeleteUserController) Execute(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.useCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
