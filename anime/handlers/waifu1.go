package handlers

import (
	"jikan-consumer/pkg/handlers"
	"net/http"
)

func Waifu1(response http.ResponseWriter, request *http.Request) {
	data := map[string]interface{}{
		"yamete": "kudasai",
		"oni":    "-chan",
	}

	handlers.WriteResponse(response, 201, data)
}
