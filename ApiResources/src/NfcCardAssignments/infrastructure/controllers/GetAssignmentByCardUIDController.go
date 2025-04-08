package controllers

import (
    "api_resources/src/NfcCardAssignments/application"
    "github.com/gin-gonic/gin"
    "net/http"
)

type GetAssignmentByCardUIDController struct {
    useCase *application.GetAssignmentByCardUID
}

func NewGetAssignmentByCardUIDController(useCase *application.GetAssignmentByCardUID) *GetAssignmentByCardUIDController {
    return &GetAssignmentByCardUIDController{useCase: useCase}
}

// GetAssignmentByCardUID godoc
// @Summary Gets an NFC card assignment by card UID
// @Description Retrieves the active assignment for a specific NFC card
// @Tags NFC Card Assignments
// @Produce json
// @Param cardUid path string true "Card UID"
// @Success 200 {object} entities.NfcCardAssignment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /api/nfc-assignments/card/{cardUid} [get]
func (c *GetAssignmentByCardUIDController) Handle(ctx *gin.Context) {
    cardUID := ctx.Param("cardUid")
    if cardUID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Card UID is required"})
        return
    }

    assignment, err := c.useCase.Execute(cardUID)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, assignment)
}