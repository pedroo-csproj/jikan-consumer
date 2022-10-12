package main

import (
	"jikan-consumer/api/server"
	"jikan-consumer/config/env"
)

func main() {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	server := server.Start()

	server.Run()
}
