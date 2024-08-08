package main

import (
	"fmt"
	"websockets/cmd/server"
)

func main() {
	server.SetupRouter().Run(":3124")
	fmt.Print("Server Started")

}
