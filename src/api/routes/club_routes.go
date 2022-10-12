package routes

import (
	"jikan-consumer/src/api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func loadClubRoutes(router *mux.Router) {
	clubRoutes := router.PathPrefix("/clubs").Subrouter()

	clubRoutes.HandleFunc("/{id}", controllers.GetClubById).Methods(http.MethodGet)
}
