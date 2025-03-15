package adapters

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/M1keTrike/EventDriven/internal/core"
	"github.com/M1keTrike/EventDriven/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type WebSocketAdapter struct {
	mu       sync.Mutex
	clients  map[string]map[*websocket.Conn]bool
	emitters map[string]*websocket.Conn
	service  *core.MessageService
}

func NewWebSocketAdapter(service *core.MessageService) *WebSocketAdapter {
	return &WebSocketAdapter{
		clients:  make(map[string]map[*websocket.Conn]bool),
		emitters: make(map[string]*websocket.Conn),
		service:  service,
	}
}

func (ws *WebSocketAdapter) HandleWebSocket(c *gin.Context) {
	topic := c.Query("topic")
	isEmitter := c.Query("emitter") == "true"

	if topic == "" {
		fmt.Println("Error: No se proporcionó un tema en la conexión WebSocket")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere un tema en la URL"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error al conectar WebSocket:", err)
		return
	}

	defer func() {
		conn.Close()
		ws.removeClient(topic, conn, isEmitter)
	}()

	ws.addClient(topic, conn, isEmitter)

	fmt.Printf("Cliente %s conectado al tema: %s\n", ws.getConnectionType(isEmitter), topic)

	for {
		var msg models.Message
		if err := conn.ReadJSON(&msg); err != nil {
			fmt.Printf("%s desconectado del tema %s\n", ws.getConnectionType(isEmitter), topic)
			break
		}

		fmt.Printf("Mensaje recibido en el servidor [%s]: %s\n", topic, msg.Content)

		ws.SendMessage(topic, &msg)
	}
}

func (ws *WebSocketAdapter) addClient(topic string, conn *websocket.Conn, isEmitter bool) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if isEmitter {

		if existingEmitter, exists := ws.emitters[topic]; exists {
			existingEmitter.Close()
		}
		ws.emitters[topic] = conn
	} else {
		if _, exists := ws.clients[topic]; !exists {
			ws.clients[topic] = make(map[*websocket.Conn]bool)
		}
		ws.clients[topic][conn] = true
	}
}

func (ws *WebSocketAdapter) removeClient(topic string, conn *websocket.Conn, isEmitter bool) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	if isEmitter {

		if ws.emitters[topic] == conn {
			delete(ws.emitters, topic)
			delete(ws.clients, topic)
			fmt.Printf("Tema %s eliminado porque el emisor se desconectó\n", topic)
		}
	} else {
		if _, exists := ws.clients[topic]; exists {
			delete(ws.clients[topic], conn)
		}
	}
}

func (ws *WebSocketAdapter) SendMessage(topic string, msg *models.Message) {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	fmt.Printf("Intentando enviar mensaje en el tema: %s\n", topic)

	if subscribers, exists := ws.clients[topic]; exists {
		for conn := range subscribers {
			fmt.Printf("Enviando mensaje a suscriptor en %s\n", topic)

			if err := conn.WriteJSON(msg); err != nil {
				fmt.Printf("Error enviando mensaje a %s: %v\n", topic, err)
				conn.Close()
				delete(subscribers, conn)
			}
		}
	} else {
		fmt.Printf("No hay suscriptores en el tema %s\n", topic)
	}
}

func (ws *WebSocketAdapter) getConnectionType(isEmitter bool) string {
	if isEmitter {
		return "Emisor"
	}
	return "Suscriptor"
}
