package controllers

import (
    "api_resources/src/NfcCardAssignments/application"
    "api_resources/src/NfcCardAssignments/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
)

type CreateAssignmentController struct {
    useCase *application.CreateAssignment
}

func NewCreateAssignmentController(useCase *application.CreateAssignment) *CreateAssignmentController {
    return &CreateAssignmentController{useCase: useCase}
}

// CreateAssignment godoc
// @Summary Creates a new NFC card assignment
// @Description Assigns an NFC card to a user
// @Tags NFC Card Assignments
// @Accept json
// @Produce json
// @Param assignment body entities.NfcCardAssignment true "Assignment Data"
// @Success 201 {object} entities.NfcCardAssignment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/nfc-assignments [post]
func (c *CreateAssignmentController) Handle(ctx *gin.Context) {
    var assignment entities.NfcCardAssignment
    if err := ctx.ShouldBindJSON(&assignment); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.useCase.Execute(&assignment); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, assignment)
}