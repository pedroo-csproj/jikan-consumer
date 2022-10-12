package server

import (
	"fmt"
	"jikan-consumer/api/routes"
	"jikan-consumer/config/env"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func StartServer() Server {
	server := Server{router: mux.NewRouter()}

	router := server.router.Host(env.Host).Subrouter()

	routes.LoadAnimeRoutes(router)
	routes.LoadClubRoutes(router)

	return server
}

func (server *Server) Run() {
	fmt.Printf("application listening on %s:%s", env.Host, env.Port)

	panic(http.ListenAndServe(":"+env.Port, server.router))
}
