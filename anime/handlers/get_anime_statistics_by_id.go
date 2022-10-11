package handlers

import (
	"jikan-consumer/pkg/handlers"
	"jikan-consumer/pkg/jikan"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAnimeStatisticsById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	anime, err := jikan.GetAnimeStatisticsById(id)
	if err != nil {
		handlers.WriteResponse(response, 404, nil)
	}

	handlers.WriteResponse(response, 200, anime)
}
