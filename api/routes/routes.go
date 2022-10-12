package routes

import (
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router) {
	loadAnimeRoutes(router)
	loadClubRoutes(router)
}
