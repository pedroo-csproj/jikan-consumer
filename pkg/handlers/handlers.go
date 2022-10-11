package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(response http.ResponseWriter, statusCode int, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	stringfiedData, _ := json.Marshal(data)

	response.Write(stringfiedData)
}
