package controllers

import (
	"github.com/M1keTrike/EventDriven/messages_aglomeration/application"
	"github.com/gin-gonic/gin"
)

type ReceiveAglomerationController struct {
	ReceiveAglomerationUseCase application.ReceiveAglomerationUseCase
}

type ReceiveAglomerationRequest struct {
	Message string `json:"message"`
}

func NewRecieveAglomerationController(ReceiveAglomerationUseCase application.ReceiveAglomerationUseCase) *ReceiveAglomerationController {
	return &ReceiveAglomerationController{ReceiveAglomerationUseCase: ReceiveAglomerationUseCase}
}

func (r *ReceiveAglomerationController) Execute(c *gin.Context) {
	var req ReceiveAglomerationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := r.ReceiveAglomerationUseCase.Execute(req.Message)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Message received successfully"})
}
