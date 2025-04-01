package controllers

import (
    "api_resources/src/NfcCardAssignments/application"
    "github.com/gin-gonic/gin"
    "net/http"
)

type GetAllAssignmentsController struct {
    useCase *application.GetAllAssignments
}

func NewGetAllAssignmentsController(useCase *application.GetAllAssignments) *GetAllAssignmentsController {
    return &GetAllAssignmentsController{useCase: useCase}
}

// GetAllAssignments godoc
// @Summary Gets all NFC card assignments
// @Description Retrieves all NFC card assignments in the system
// @Tags NFC Card Assignments
// @Produce json
// @Success 200 {array} entities.NfcCardAssignment
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/nfc-assignments [get]
func (c *GetAllAssignmentsController) Handle(ctx *gin.Context) {
    assignments, err := c.useCase.Execute()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, assignments)
}