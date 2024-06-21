package handlers

import (
	"fmt"
	"net/http"
	"websockets/internals/hub"
	"websockets/internals/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()
	hub.Clients[conn] = true
	for {
		var msg models.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error reading JSON:", err)
			delete(hub.Clients, conn)
			return
		}
		hub.Broadcast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-hub.Broadcast
		for client := range hub.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				fmt.Println("Error writing JSON:", err)
				client.Close()
				delete(hub.Clients, client)
			}
		}
	}
}
