package controllers

import (
    "api_resources/src/NfcCardAssignments/application"
    "api_resources/src/NfcCardAssignments/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
)

type UpdateAssignmentController struct {
    useCase *application.UpdateAssignment
}

func NewUpdateAssignmentController(useCase *application.UpdateAssignment) *UpdateAssignmentController {
    return &UpdateAssignmentController{useCase: useCase}
}

// UpdateAssignment godoc
// @Summary Updates an NFC card assignment
// @Description Updates the details of an existing NFC card assignment
// @Tags NFC Card Assignments
// @Accept json
// @Produce json
// @Param id path int true "Assignment ID"
// @Param assignment body entities.NfcCardAssignment true "Updated Assignment Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/nfc-assignments/{id} [put]
func (c *UpdateAssignmentController) Handle(ctx *gin.Context) {
    var assignment entities.NfcCardAssignment
    if err := ctx.ShouldBindJSON(&assignment); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.useCase.Execute(&assignment); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.Status(http.StatusOK)
}