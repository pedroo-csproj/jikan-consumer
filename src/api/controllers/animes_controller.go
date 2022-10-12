package controllers

import (
	services "jikan-consumer/src/services/anime"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAnimeById(response http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]

	getAnimeResult := services.GetAnimeById(id)
	if !getAnimeResult.Success {
		writeResponse(response, getAnimeResult.StatusCode, getAnimeResult.Errors)
	}

	writeResponse(response, getAnimeResult.StatusCode, getAnimeResult.Data)
}

func GetAnimeStatisticsById(response http.ResponseWriter, request *http.Request) {
	id := mux.Vars(request)["id"]

	getAnimeStatisticsResult := services.GetAnimeStatisticsById(id)
	if !getAnimeStatisticsResult.Success {
		writeResponse(response, getAnimeStatisticsResult.StatusCode, getAnimeStatisticsResult.Errors)
	}

	writeResponse(response, getAnimeStatisticsResult.StatusCode, getAnimeStatisticsResult.Data)
}
