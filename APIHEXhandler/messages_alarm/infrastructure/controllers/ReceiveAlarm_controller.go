package controllers

import (
    "github.com/M1keTrike/EventDriven/messages_alarm/application"
    "github.com/gin-gonic/gin"
)

type ReceiveAlarmController struct {
    ReceiveAlarmUseCase application.ReceiveAlarmUseCase
}

type ReceiveAlarmRequest struct {
    Message string `json:"message"`
}

func NewReceiveAlarmController(ReceiveAlarmUseCase application.ReceiveAlarmUseCase) *ReceiveAlarmController {
    return &ReceiveAlarmController{ReceiveAlarmUseCase: ReceiveAlarmUseCase}
}

func (controller *ReceiveAlarmController) Handle(c *gin.Context) {
    var request ReceiveAlarmRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err := controller.ReceiveAlarmUseCase.Execute(request.Message)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"status": "success"})
}