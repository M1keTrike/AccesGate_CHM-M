package controllers

import (
	"github.com/M1keTrike/EventDriven/messages_nfc/application"
	"github.com/gin-gonic/gin"
)

type ReceiveNFCMessageController struct {
	rm_uc application.RecieveNFCMessageUseCase
}

type ReceiveMessageRequest struct {
	Uid string `json:"uid"`
}

func NewReceiveMessageNFCController(rm_uc application.RecieveNFCMessageUseCase) *ReceiveNFCMessageController {
	return &ReceiveNFCMessageController{rm_uc: rm_uc}
}

func (r *ReceiveNFCMessageController) Execute(c *gin.Context) {
	var req ReceiveMessageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	err := r.rm_uc.Execute(req.Uid)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Message received successfully"})
}
