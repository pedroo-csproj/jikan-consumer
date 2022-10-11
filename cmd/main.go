package main

import (
	"jikan-consumer/pkg/env"
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
