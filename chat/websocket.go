package chat

import (
	"log"
	"net/http"
	"startup-chatrooms/db"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for testing)
	},
}

// HandleWebSocket upgrades HTTP to WebSocket and manages incoming messages
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer ws.Close()

	// Continuously read messages from the WebSocket connection
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// Log the received message
		log.Printf("Received: %s", message)

		// Save the message to DynamoDB
		err = db.SaveMessageToDynamoDB(string(message), "startup-idea-123") // Example partition key
		if err != nil {
			log.Println("DynamoDB save error:", err)
			break
		}

		// Echo the message back to the client
		err = ws.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
