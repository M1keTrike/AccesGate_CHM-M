package controllers

import (
	"github.com/M1keTrike/EventDriven/messages_fingerprint/application"
	"github.com/gin-gonic/gin"
)

type ReceiveFingerprintController struct {
	ReceiveFingerprintUseCase application.ReceiveFingerprintUseCase
}

type ReceiveFingerprintRequest struct {
	Message string `json:"message"`
}

func NewReceiveFingerprintController(ReceiveFingerprintUseCase application.ReceiveFingerprintUseCase) *ReceiveFingerprintController {
	return &ReceiveFingerprintController{ReceiveFingerprintUseCase: ReceiveFingerprintUseCase}
}

func (controller *ReceiveFingerprintController) Handle(c *gin.Context) {
	var request ReceiveFingerprintRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := controller.ReceiveFingerprintUseCase.Execute(request.Message)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}