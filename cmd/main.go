package main

import (
	"fmt"
	"websockets/cmd/server"
)

func main() {
	server.SetupRouter().Run(":8080")
	fmt.Print("Server Started")

}
