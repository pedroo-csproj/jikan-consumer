package routes

import (
	"jikan-consumer/api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func loadClubRoutes(router *mux.Router) {
	clubRoutes := router.PathPrefix("/clubs").Subrouter()

	clubRoutes.HandleFunc("/{id}", controllers.GetClubById).Methods(http.MethodGet)
}
