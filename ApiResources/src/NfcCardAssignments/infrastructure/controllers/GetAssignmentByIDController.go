package controllers

import (
    "api_resources/src/NfcCardAssignments/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type GetAssignmentByIDController struct {
    useCase *application.GetAssignmentByID
}

func NewGetAssignmentByIDController(useCase *application.GetAssignmentByID) *GetAssignmentByIDController {
    return &GetAssignmentByIDController{useCase: useCase}
}

// GetAssignmentByID godoc
// @Summary Gets an NFC card assignment by ID
// @Description Retrieves a specific NFC card assignment using its ID
// @Tags NFC Card Assignments
// @Produce json
// @Param id path int true "Assignment ID"
// @Success 200 {object} entities.NfcCardAssignment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/nfc-assignments/{id} [get]
func (c *GetAssignmentByIDController) Handle(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    assignment, err := c.useCase.Execute(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, assignment)
}