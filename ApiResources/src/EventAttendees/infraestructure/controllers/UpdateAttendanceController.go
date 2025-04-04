package controllers

import (
    "api_resources/src/EventAttendees/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type UpdateAttendanceController struct {
    useCase *application.UpdateAttendanceStatus
}

func NewUpdateAttendanceController(useCase *application.UpdateAttendanceStatus) *UpdateAttendanceController {
    return &UpdateAttendanceController{useCase: useCase}
}

type UpdateAttendanceRequest struct {
    Attended bool `json:"attended"`
}

func (c *UpdateAttendanceController) Execute(ctx *gin.Context) {
    eventID, err := strconv.Atoi(ctx.Param("eventId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
        return
    }

    userID, err := strconv.Atoi(ctx.Param("userId"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var req UpdateAttendanceRequest
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }

    err = c.useCase.Execute(eventID, userID, req.Attended)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Attendance status updated successfully"})
}