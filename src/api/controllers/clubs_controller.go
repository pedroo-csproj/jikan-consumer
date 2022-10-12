package controllers

import (
	services "jikan-consumer/src/services/club"
	"net/http"

	"github.com/gorilla/mux"
)

func GetClubById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	club, _ := services.GetClubById(id)

	if club.ID == 0 {
		writeResponse(response, 404, nil)
	}

	writeResponse(response, 200, club)
}
