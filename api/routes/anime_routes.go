package routes

import (
	"jikan-consumer/anime/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadAnimeRoutes(router *mux.Router) {
	animeRoutes := router.PathPrefix("/animes").Subrouter()

	animeRoutes.HandleFunc("/{id}", handlers.GetAnimeById).Methods(http.MethodGet)
	animeRoutes.HandleFunc("/{id}/statistics", handlers.GetAnimeStatisticsById).Methods(http.MethodGet)
}
