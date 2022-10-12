package routes

import (
	"jikan-consumer/clubs/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadClubRoutes(router *mux.Router) {
	clubRoutes := router.PathPrefix("/clubs").Subrouter()

	clubRoutes.HandleFunc("/{id}", handlers.GetClubById).Methods(http.MethodGet)
}
