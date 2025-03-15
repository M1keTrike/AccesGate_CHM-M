package controllers

import (
	"github.com/M1keTrike/EventDriven/messages_v_access/application"
	"github.com/gin-gonic/gin"
)

type ReceiveVAccessController struct {
	ReceiveVAccessUseCase application.ReceiveVAccessUseCase
}

type ReceiveVAccessRequest struct {
	Message string `json:"message"`
}

func NewRecieveVAccessController(ReceiveVAccessUseCase application.ReceiveVAccessUseCase) *ReceiveVAccessController {
	return &ReceiveVAccessController{ReceiveVAccessUseCase: ReceiveVAccessUseCase}
}

func (r *ReceiveVAccessController) Execute(c *gin.Context) {
	var req ReceiveVAccessRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := r.ReceiveVAccessUseCase.Execute(req.Message)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Message received successfully"})
}
