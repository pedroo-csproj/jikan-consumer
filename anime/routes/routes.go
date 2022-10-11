package anime

import (
	"jikan-consumer/anime/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router) {
	animeRoutes := router.PathPrefix("/animes").Subrouter()

	animeRoutes.HandleFunc("/{id}", handlers.GetAnimeById).Methods(http.MethodGet)
	animeRoutes.HandleFunc("/waifu22", handlers.Waifu2).Methods(http.MethodGet)
}
