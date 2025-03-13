package server

import (
	"github.com/M1keTrike/EventDriven/internal/adapters"
	"github.com/M1keTrike/EventDriven/internal/core"
	"github.com/gin-gonic/gin"
)

func StartServer() {

	r := gin.Default()

	repo := adapters.NewInMemoryRepository()
	service := core.NewMessageService(repo)
	wsAdapter := adapters.NewWebSocketAdapter(service)

	r.GET("/ws", wsAdapter.HandleWebSocket)

	r.Run(":8081")
}
