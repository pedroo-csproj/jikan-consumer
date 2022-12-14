package controllers

import (
	services "jikan-consumer/src/services/club"
	"net/http"

	"github.com/gorilla/mux"
)

func GetClubById(response http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]

	getClubResult := services.GetClubById(id)
	if !getClubResult.Success {
		writeResponse(response, getClubResult.StatusCode, getClubResult.Errors)
	}

	writeResponse(response, getClubResult.StatusCode, getClubResult.Data)
}
