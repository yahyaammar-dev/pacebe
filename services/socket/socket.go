package socket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for testing; adjust for production
	},
}

type Handler struct{}

func NewConnection() *Handler {
	return &Handler{}
}

func (h *Handler) Run() error {
	http.HandleFunc("/", h.handleConnection)
	fmt.Println("WebSocket server is listening on port 9999...")
	return http.ListenAndServe(":9999", nil)
}

func (h *Handler) handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	for {
		// Read message from client
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// Print the received message
		fmt.Printf("Received: %d\n", messageType)
		fmt.Printf("Received: %s\n", msg)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, msg); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}

	fmt.Println("Client disconnected")
}
