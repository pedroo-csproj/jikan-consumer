package handlers

import (
	"jikan-consumer/pkg/handlers"
	"jikan-consumer/pkg/jikan"
	"net/http"

	"github.com/gorilla/mux"
)

func GetClubById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	club, _ := jikan.GetClubById(id)

	if club.ID == 0 {
		handlers.WriteResponse(response, 404, nil)
	}

	handlers.WriteResponse(response, 200, club)
}
