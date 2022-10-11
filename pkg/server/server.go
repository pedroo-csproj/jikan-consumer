package server

import (
	"fmt"
	anime "jikan-consumer/anime/routes"
	character "jikan-consumer/character/routes"
	"jikan-consumer/pkg/env"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func StartServer() Server {
	server := Server{router: mux.NewRouter()}

	router := server.router.Host(env.Host).Subrouter()

	anime.LoadRoutes(router)
	character.LoadRoutes(router)

	return server
}

func (server *Server) Run() {
	fmt.Printf("application listening on %s:%s", env.Host, env.Port)

	panic(http.ListenAndServe(":"+env.Port, server.router))
}
