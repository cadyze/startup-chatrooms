package main

import (
	"log"
	"net/http"
	"startup-chatrooms/chat"
)

func main() {
	// Register WebSocket handler
	http.HandleFunc("/ws", chat.HandleWebSocket)

	log.Println("WebSocket server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
