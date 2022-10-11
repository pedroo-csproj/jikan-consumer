package character

import (
	"jikan-consumer/character/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router) {
	characterRoutes := router.PathPrefix("/animes").Subrouter()

	characterRoutes.HandleFunc("/character1", handlers.Character1).Methods(http.MethodGet)
}
