package handlers

import (
	"fmt"
	"net/http"
	"websockets/internals/hub"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func CreateRoomId(roomId string, roomName string) error {
	hub.Mu.Lock()
	defer hub.Mu.Unlock()

	if _, ok := hub.Rooms[roomId]; ok {
		return fmt.Errorf("room already exists")
	}
	hub.Rooms[roomId] = make(map[*websocket.Conn]bool)
	hub.RoomInfo[roomId] = roomName
	return nil
}

func CreateRoomHandlers(ctx *gin.Context) {
	roomId := ctx.Param("roomId")
	roomName := ctx.Param("roomName")

	// Check if roomId is provided
	if roomId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "roomId parameter is required"})
		return
	}

	fmt.Printf("Received request to create room with ID: %s\n", roomId)
	err := CreateRoomId(roomId, roomName)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Room created successfully", "roomId": roomId, "roomName": roomName})
}
