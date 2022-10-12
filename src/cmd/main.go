package main

import (
	"jikan-consumer/src/api/server"
	"jikan-consumer/src/config/env"
)

func main() {
	err := env.Load(".env")

	if err != nil {
		panic(err)
	}

	server := server.Start()

	server.Run()
}
