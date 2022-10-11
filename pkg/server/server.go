package server

import (
	"fmt"
	anime "jikan-consumer/anime/routes"
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

	return server
}

func (server *Server) Run() {
	http.ListenAndServe(":"+"4200", server.router)

	fmt.Println("application listening on port '4200'")
}
