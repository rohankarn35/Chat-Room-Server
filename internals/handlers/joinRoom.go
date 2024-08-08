package handlers

import (
	"net/http"
	"websockets/internals/hub"

	"github.com/gin-gonic/gin"
)

func JoinRoom(ctx *gin.Context) {
	roomId := ctx.Query("roomId")
	if roomId == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Error"})
		return
	}
	if _, ok := hub.Rooms[roomId]; !ok {
		ctx.JSON(http.StatusConflict, gin.H{"Error": "Room Does Not Exist"})
		return
	}

	roomName := hub.RoomInfo[roomId]

	ctx.JSON(http.StatusAccepted, gin.H{"roomId": roomId, "roomName": roomName})

}

// mainAxisAlignment: MainAxisAlignment.center,
