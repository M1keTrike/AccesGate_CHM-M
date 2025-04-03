package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type WebSocketFingerprintEmitter struct {
	conn  *websocket.Conn
	topic string
}

func NewWebSocketFingerprintEmitter(webSocketURL string) (*WebSocketFingerprintEmitter, error) {
	wsEndpoint := fmt.Sprintf("%s/ws?topic=fingerprint&emitter=true", webSocketURL)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Printf("Error connecting to WebSocket Server: %v", err)
		return nil, err
	}

	return &WebSocketFingerprintEmitter{
		conn:  conn,
		topic: "fingerprint",
	}, nil
}

func (w *WebSocketFingerprintEmitter) Send(content []byte) error {
	message := map[string]string{"content": string(content)}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = w.conn.WriteMessage(websocket.TextMessage, jsonMessage)
	if err != nil {
		log.Printf("Error sending message to WebSocket: %v", err)
		return err
	}

	fmt.Printf("Message sent to WebSocket [%s]: %s\n", w.topic, content)
	return nil
}

func (w *WebSocketFingerprintEmitter) Close() {
	if w.conn != nil {
		w.conn.Close()
		fmt.Println("WebSocket connection closed")
	}
}