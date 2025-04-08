package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/M1keTrike/EventDriven/messages_nfc/application"
	"github.com/gin-gonic/gin"
)

type ReceiveNFCController struct {
	ReceiveNFCUseCase application.RecieveNFCMessageUseCase
}

type ReceiveNFCRequest struct {
	Message string `json:"message"`
}

func NewReceiveNFCController(ReceiveNFCUseCase application.RecieveNFCMessageUseCase) *ReceiveNFCController {
	return &ReceiveNFCController{ReceiveNFCUseCase: ReceiveNFCUseCase}
}

func (controller *ReceiveNFCController) Handle(c *gin.Context) {
	var request ReceiveNFCRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Printf("Error parsing request: %v\n", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	
	var messageData map[string]interface{}
	if err := json.Unmarshal([]byte(request.Message), &messageData); err != nil {
		fmt.Printf("Error parsing message content: %v\n", err)
		c.JSON(400, gin.H{"error": "Invalid message format"})
		return
	}

	
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

	// Create clean message structure
	cleanMessage := map[string]string{
		"mac":    mac,
		"evento": evento,
	}

	// Convert to JSON
	finalMessage, err := json.Marshal(cleanMessage)
	if err != nil {
		fmt.Printf("NFC Controller - Error creating message: %v\n", err)
		c.JSON(500, gin.H{"error": "Error creating message"})
		return
	}

	err = controller.ReceiveNFCUseCase.Execute(string(finalMessage))
	if err != nil {
		fmt.Printf("NFC Controller - Error executing use case: %v\n", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}
