package hub

import (
	"websockets/internals/models"

	"github.com/gorilla/websocket"
)

var Clients = make(map[*websocket.Conn]bool)
var Broadcast = make(chan models.Message)
