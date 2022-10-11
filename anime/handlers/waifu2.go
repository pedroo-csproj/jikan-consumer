package handlers

import (
	"jikan-consumer/pkg/handlers"
	"net/http"
)

func Waifu2(response http.ResponseWriter, request *http.Request) {
	data := map[string]interface{}{
		"yamete": "kudasai11",
		"oni":    "-chan11",
	}

	handlers.WriteResponse(response, 200, data)
}
