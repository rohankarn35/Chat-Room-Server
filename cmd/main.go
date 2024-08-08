package main

import (
	"websockets/cmd/server"
)

func main() {
	server.SetupRouter().Run(":3124")

}
