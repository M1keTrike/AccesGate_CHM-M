package controllers

import (
    "api_resources/src/NfcCardAssignments/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type DeactivateAssignmentController struct {
    useCase *application.DeactivateAssignment
}

func NewDeactivateAssignmentController(useCase *application.DeactivateAssignment) *DeactivateAssignmentController {
    return &DeactivateAssignmentController{useCase: useCase}
}

// DeactivateAssignment godoc
// @Summary Deactivates an NFC card assignment
// @Description Sets an NFC card assignment as inactive
// @Tags NFC Card Assignments
// @Produce json
// @Param id path int true "Assignment ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/nfc-assignments/{id}/deactivate [put]
func (c *DeactivateAssignmentController) Handle(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
        return
    }

    if err := c.useCase.Execute(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusOK)
}