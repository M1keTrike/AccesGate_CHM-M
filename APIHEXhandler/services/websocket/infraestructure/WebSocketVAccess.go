package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type WebSocketVAccessEmitter struct {
	conn  *websocket.Conn
	topic string
}

func NewWebSocketVAccessEmitter(webSocketURL string) (*WebSocketVAccessEmitter, error) {
	wsEndpoint := fmt.Sprintf("%s/ws?topic=vaccess&emitter=true", webSocketURL)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Printf("Error al conectar con WebSocket Server: %v", err)
		return nil, err
	}

	fmt.Printf("Conexión establecida con WebSocket Server [%s]\n", wsEndpoint)

	return &WebSocketVAccessEmitter{
		conn:  conn,
		topic: "vaccess",
	}, nil
}

func (w *WebSocketVAccessEmitter) Send(content []byte) error {
	message := map[string]string{"content": string(content)}
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = w.conn.WriteMessage(websocket.TextMessage, jsonMessage)
	if err != nil {
		log.Printf("Error enviando mensaje al WebSocket: %v", err)
		return err
	}

	fmt.Printf("Mensaje enviado al WebSocket [%s]: %s\n", w.topic, content)
	return nil
}

func (w *WebSocketVAccessEmitter) Close() {
	if w.conn != nil {
		w.conn.Close()
		fmt.Println("Conexión WebSocket cerrada")
	}
}
