package routes

import (
	"jikan-consumer/src/api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func loadAnimeRoutes(router *mux.Router) {
	animeRoutes := router.PathPrefix("/animes").Subrouter()

	animeRoutes.HandleFunc("/{id}", controllers.GetAnimeById).Methods(http.MethodGet)
	animeRoutes.HandleFunc("/{id}/statistics", controllers.GetAnimeStatisticsById).Methods(http.MethodGet)
}
