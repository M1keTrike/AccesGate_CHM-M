package controllers

import (
    "github.com/M1keTrike/EventDriven/messages_fingerprint_registration/application"
    "github.com/gin-gonic/gin"
)

type ReceiveFingerprintRegistrationController struct {
    ReceiveFingerprintRegistrationUseCase application.ReceiveFingerprintRegistrationUseCase
}

type ReceiveFingerprintRegistrationRequest struct {
    Message string `json:"message"`
}

func NewReceiveFingerprintRegistrationController(ReceiveFingerprintRegistrationUseCase application.ReceiveFingerprintRegistrationUseCase) *ReceiveFingerprintRegistrationController {
    return &ReceiveFingerprintRegistrationController{ReceiveFingerprintRegistrationUseCase: ReceiveFingerprintRegistrationUseCase}
}

func (controller *ReceiveFingerprintRegistrationController) Handle(c *gin.Context) {
    var request ReceiveFingerprintRegistrationRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    err := controller.ReceiveFingerprintRegistrationUseCase.Execute(request.Message)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"status": "success"})
}