package infrastructure

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/gorilla/websocket"
)

type WebSocketFingerprintRegistrationEmitter struct {
    conn  *websocket.Conn
    topic string
}

func NewWebSocketFingerprintRegistrationEmitter(webSocketURL string) (*WebSocketFingerprintRegistrationEmitter, error) {
    wsEndpoint := fmt.Sprintf("%s/ws?topic=fingerprint_registration&emitter=true", webSocketURL)

    conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
    if err != nil {
        log.Printf("Error connecting to WebSocket Server: %v", err)
        return nil, err
    }

    return &WebSocketFingerprintRegistrationEmitter{
        conn:  conn,
        topic: "fingerprint_registration",
    }, nil
}

func (w *WebSocketFingerprintRegistrationEmitter) Send(content []byte) error {
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

func (w *WebSocketFingerprintRegistrationEmitter) Close() {
    if w.conn != nil {
        w.conn.Close()
        fmt.Println("WebSocket connection closed")
    }
}