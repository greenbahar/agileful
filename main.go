package main

import "testTask/cmd/server"

func main() {
	server:=server.NewAppServer()
	server.InitObjects()
	server.SetupRouter()
	server.Run()
}



