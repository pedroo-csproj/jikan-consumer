package jikan

import (
	"encoding/json"
	"fmt"
	"io"
	"jikan-consumer/pkg/env"
	"jikan-consumer/pkg/jikan/models"
	"net/http"
)

func GetAnimeById(id string) (models.GetAnimeById, error) {
	resp, err := http.Get(fmt.Sprintf("%sanime/%s/full", env.JikanApi, id))
	if err != nil {
		return models.GetAnimeById{}, err
	}

	if resp.Status == "404" {
		return models.GetAnimeById{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.GetAnimeById{}, err
	}

	responseWrapper := models.ResponseWrapper[models.GetAnimeById]{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return models.GetAnimeById{}, err
	}

	return responseWrapper.Data, nil
}
