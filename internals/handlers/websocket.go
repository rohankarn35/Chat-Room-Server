package handlers

import (
	"log"
	"websockets/internals/hub"
	"websockets/internals/models"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func HandleConnections(ctx *gin.Context) {
	writer := ctx.Writer
	requester := ctx.Request

	roomID := ctx.Param("roomId")

	// Upgrade the HTTP request to a WebSocket connection
	ws, err := hub.Upgrader.Upgrade(writer, requester, nil)
	if err != nil {
		log.Println("Error occurred in websocket upgrader", err.Error())
		return
	}
	defer ws.Close()

	// Lock and handle the room creation
	hub.Mu.Lock()
	_, ok := hub.Rooms[roomID]
	hub.Mu.Unlock()

	if !ok {
		errMsg := websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Room not found, create room first")
		ws.WriteMessage(websocket.CloseMessage, errMsg)
		return
	}

	hub.Mu.Lock()
	hub.Rooms[roomID][ws] = true
	hub.Mu.Unlock()

	// Listen for incoming messages
	for {
		var msg models.Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("Error occurred while reading the message", err.Error())
			hub.Mu.Lock()
			delete(hub.Rooms[roomID], ws)
			if len(hub.Rooms[roomID]) == 0 {
				delete(hub.Rooms, roomID)
			}
			hub.Mu.Unlock()
			break
		}

		// Broadcast the received message to the room
		hub.Broadcast <- models.RoomMessage{
			RoomID:  roomID,
			Message: msg,
			Sender:  ws,
		}
	}
}

func HandleMessages() {
	for roomMsg := range hub.Broadcast {
		hub.Mu.Lock()
		clients, ok := hub.Rooms[roomMsg.RoomID]
		hub.Mu.Unlock()
		if !ok {
			continue
		}

		for client := range clients {
			if client != roomMsg.Sender {
				err := client.WriteJSON(roomMsg.Message)
				if err != nil {
					log.Println("Error occurred while sending the message", err.Error())
					client.Close()
					hub.Mu.Lock()
					delete(clients, client)
					if len(clients) == 0 {
						delete(hub.Rooms, roomMsg.RoomID)
					}
					hub.Mu.Unlock()
				}
			}
		}
	}
}
