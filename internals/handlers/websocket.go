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
	rooms     = make(map[string]map[*websocket.Conn]bool)
	broadcast = make(chan models.RoomMessage)
	mu        sync.Mutex
)

func HandleConnections(ctx *gin.Context) {
	writer := ctx.Writer
	requester := ctx.Request

	roomID := ctx.Query("room_id")

	ws, err := upgrader.Upgrade(writer, requester, nil)
	if err != nil {
		log.Println("Error occurred in websocket upgrader", err.Error())
		return
	}
	defer ws.Close()

	mu.Lock()
	if _, ok := rooms[roomID]; !ok {
		rooms[roomID] = make(map[*websocket.Conn]bool)
	}
	rooms[roomID][ws] = true
	mu.Unlock()

	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("Error occurred while reading the msg", err.Error())
			mu.Lock()
			delete(rooms[roomID], ws)
			if len(rooms[roomID]) == 0 {
				delete(rooms, roomID)
			}
			mu.Unlock()
			break
		}
		broadcast <- models.RoomMessage{
			RoomID:  roomID,
			Message: msg,
			Sender:  ws,
		}
	}
}

func HandleMessage() {
	for roomMsg := range broadcast {
		mu.Lock()
		clients, ok := rooms[roomMsg.RoomID]
		mu.Unlock()
		if !ok {
			continue
		}

		for client := range clients {
			if client != roomMsg.Sender {
				err := client.WriteJSON(roomMsg.Message)
				if err != nil {
					log.Println("Error occurred while sending the msg", err.Error())
					client.Close()
					mu.Lock()
					delete(clients, client)
					if len(clients) == 0 {
						delete(rooms, roomMsg.RoomID)
					}
					mu.Unlock()
				}
			}
		}
	}
}
