package handlers

import (
	"log"
	"net/http"
	"sync"
	"websockets/internals/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan models.Message)
	mu        sync.Mutex
)

func HandleConnections(ctx *gin.Context) {
	writer := ctx.Writer
	requester := ctx.Request

	ws, err := upgrader.Upgrade(writer, requester, nil)
	if err != nil {
		log.Println("Error Occured in websocket upgrader", err.Error())
		return
	}
	defer ws.Close()
	mu.Lock()
	clients[ws] = true
	mu.Unlock()
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("Error Occured while reading the msg", err.Error())
			mu.Lock()
			delete(clients, ws)
			mu.Unlock()
			break
		}
		broadcast <- msg

	}
}

func HandleMessage() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Error Occured while sending the msg", err.Error())

				delete(clients, client)

			}

		}
		mu.Unlock()
	}
}
