package hub

import (
	"net/http"
	"sync"
	"websockets/internals/models"

	"github.com/gorilla/websocket"
)

var (
	Upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		}}
	Rooms     = make(map[string]map[*websocket.Conn]bool)
	Broadcast = make(chan models.RoomMessage)
	Mu        sync.Mutex
	RoomInfo  = make(map[string]string)
)
