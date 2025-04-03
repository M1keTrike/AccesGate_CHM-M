package controllers

import (
	"encoding/json"

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

	// Parse the incoming message
	var messageData map[string]interface{}
	if err := json.Unmarshal([]byte(req.Message), &messageData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid message format"})
		return
	}

	// Extract fields
	mac, ok := messageData["mac"].(string)
	if !ok {
		c.JSON(400, gin.H{"error": "Missing or invalid MAC address"})
		return
	}

	evento, ok := messageData["evento"].(string)
	if !ok {
		c.JSON(400, gin.H{"error": "Missing or invalid evento"})
		return
	}

	// Create clean message
	cleanMessage := map[string]string{
		"mac":    mac,
		"evento": evento,
	}

	// Convert to JSON
	finalMessage, err := json.Marshal(cleanMessage)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating message"})
		return
	}

	err = r.ReceiveAglomerationUseCase.Execute(string(finalMessage))
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Message received successfully"})
}
