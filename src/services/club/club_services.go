package services

import (
	"encoding/json"
	"fmt"
	"io"
	"jikan-consumer/src/config/env"
	"jikan-consumer/src/models"
	"jikan-consumer/src/services/club/dtos"
	"net/http"
)

func GetClubById(id string) (dtos.GetClubById, error) {
	resp, err := http.Get(fmt.Sprintf("%sclubs/%s", env.JikanApi, id))
	if err != nil {
		return dtos.GetClubById{}, err
	}

	if resp.Status == "404" {
		return dtos.GetClubById{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.GetClubById{}, err
	}

	responseWrapper := models.ResponseWrapper[dtos.GetClubById]{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return dtos.GetClubById{}, err
	}

	return responseWrapper.Data, nil
}
