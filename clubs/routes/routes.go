package routes

import (
	"jikan-consumer/clubs/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router) {
	animeRoutes := router.PathPrefix("/clubs").Subrouter()

	animeRoutes.HandleFunc("/{id}", handlers.GetClubById).Methods(http.MethodGet)
}
