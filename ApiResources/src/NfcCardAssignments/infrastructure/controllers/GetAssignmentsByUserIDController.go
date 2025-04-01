package controllers

import (
    "api_resources/src/NfcCardAssignments/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type GetAssignmentsByUserIDController struct {
    useCase *application.GetAssignmentsByUserID
}

func NewGetAssignmentsByUserIDController(useCase *application.GetAssignmentsByUserID) *GetAssignmentsByUserIDController {
    return &GetAssignmentsByUserIDController{useCase: useCase}
}

// GetAssignmentsByUserID godoc
// @Summary Gets all NFC card assignments for a user
// @Description Retrieves all NFC card assignments associated with a specific user
// @Tags NFC Card Assignments
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} entities.NfcCardAssignment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/nfc-assignments/user/{userId} [get]
func (c *GetAssignmentsByUserIDController) Handle(ctx *gin.Context) {
    userID, err := strconv.Atoi(ctx.Param("userId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
        return
    }

    assignments, err := c.useCase.Execute(userID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, assignments)
}