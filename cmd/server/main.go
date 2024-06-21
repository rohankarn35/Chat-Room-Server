package main

import (
	"fmt"
	"net/http"
	"websockets/internals/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Homepage)
	http.HandleFunc("/ws", handlers.HandleConnections)
	go handlers.HandleMessages()
	fmt.Println("Server started at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Error occurred in server starting: " + err.Error())
	}
}
