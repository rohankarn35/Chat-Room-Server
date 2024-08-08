package server

import (
	"websockets/internals/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // Allow your Flutter web app's URL
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	router.GET("/createRoom/:roomId/:roomName", handlers.CreateRoomHandlers)

	router.GET("/join", handlers.JoinRoom)

	router.GET("/ws/:roomId", handlers.HandleConnections)
	go handlers.HandleMessages()

	return router

}
