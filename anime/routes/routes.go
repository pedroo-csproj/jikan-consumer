package anime

import (
	"jikan-consumer/anime/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router) {
	router.HandleFunc("/waifu11", handlers.Waifu1).Methods(http.MethodGet)
	router.HandleFunc("/waifu22", handlers.Waifu2).Methods(http.MethodGet)
}
