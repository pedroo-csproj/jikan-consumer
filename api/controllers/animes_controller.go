package controllers

import (
	services "jikan-consumer/services/anime"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAnimeById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	anime, _ := services.GetAnimeById(id)

	if anime.ID == 0 {
		writeResponse(response, 404, nil)
	}

	writeResponse(response, 200, anime)
}

func GetAnimeStatisticsById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	anime, err := services.GetAnimeStatisticsById(id)
	if err != nil {
		writeResponse(response, 404, nil)
	}

	writeResponse(response, 200, anime)
}
