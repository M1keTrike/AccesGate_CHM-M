package controllers

import (
    "api_resources/src/Users/application"
    "github.com/gin-gonic/gin"
    "net/http"
)

type GetUsersByRoleController struct {
    useCase application.GetUsersByRole
}

func NewGetUsersByRoleController(useCase application.GetUsersByRole) *GetUsersByRoleController {
    return &GetUsersByRoleController{useCase: useCase}
}

// GetUsersByRole godoc
// @Summary Get users by role
// @Description Retrieves all users with a specific role
// @Tags Users
// @Produce json
// @Param role path string true "User role"
// @Success 200 {array} entities.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /users/role/{role} [get]
func (c *GetUsersByRoleController) Execute(ctx *gin.Context) {
    role := ctx.Param("role")
    if role == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Role is required"})
        return
    }

    users, err := c.useCase.Execute(role)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, users)
}