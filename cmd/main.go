package main

import (
	"jikan-consumer/config/env"
	"jikan-consumer/pkg/server"
)

func main() {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	server := server.StartServer()

	server.Run()
}
