package main

import "jikan-consumer/pkg/server"

func main() {
	server := server.StartServer()

	server.Run()
}
