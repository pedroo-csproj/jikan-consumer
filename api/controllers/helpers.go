package controllers

import (
	"encoding/json"
	"net/http"
)

func writeResponse(response http.ResponseWriter, statusCode int, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)

	stringfiedData, _ := json.Marshal(data)

	response.Write(stringfiedData)
}
