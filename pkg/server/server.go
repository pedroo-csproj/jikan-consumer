package server

import (
	"fmt"
	anime "jikan-consumer/anime/routes"
	character "jikan-consumer/character/routes"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func StartServer() Server {
	server := Server{router: mux.NewRouter()}

	router := server.router.Host("localhost").Subrouter()

	anime.LoadRoutes(router)
	character.LoadRoutes(router)

	return server
}

func (server *Server) Run() {
	fmt.Println("application listening on port '4200'")

	panic(http.ListenAndServe(":"+"4200", server.router))
}
