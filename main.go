package main

import (
	"log"
	"net/http"
	"startup-chatrooms/auth"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/callback", auth.CallbackHandler)

	// Protect your routes with the AuthMiddleware
	http.Handle("/chat", auth.AuthMiddleware(http.HandlerFunc(chatHandler)))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// Your chat handler code here
	w.Write([]byte("Welcome to the chatroom!"))
}
