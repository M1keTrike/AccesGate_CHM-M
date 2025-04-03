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

func NewReceiveAglomerationController(ReceiveAglomerationUseCase application.ReceiveAglomerationUseCase) *ReceiveAglomerationController {
    return &ReceiveAglomerationController{ReceiveAglomerationUseCase: ReceiveAglomerationUseCase}
}

func (controller *ReceiveAglomerationController) Handle(c *gin.Context) {
    var request ReceiveAglomerationRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    // Parse the incoming message
    var messageData map[string]interface{}
    if err := json.Unmarshal([]byte(request.Message), &messageData); err != nil {
        c.JSON(400, gin.H{"error": "Invalid message format"})
        return
    }

    // Validate required fields
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

    // Clean and standardize the message format
    cleanMessage := map[string]string{
        "mac":    mac,
        "evento": evento,
    }
    jsonMessage, err := json.Marshal(cleanMessage)
    if err != nil {
        c.JSON(500, gin.H{"error": "Error formatting message"})
        return
    }

    err = controller.ReceiveAglomerationUseCase.Execute(string(jsonMessage))
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{"status": "success"})
}