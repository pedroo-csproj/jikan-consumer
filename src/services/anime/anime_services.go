package services

import (
	"encoding/json"
	"fmt"
	"io"
	"jikan-consumer/src/config/env"
	"jikan-consumer/src/models"
	"jikan-consumer/src/services/anime/dtos"
	"net/http"
)

func GetAnimeById(id string) (dtos.GetAnimeById, error) {
	resp, err := http.Get(fmt.Sprintf("%sanime/%s/full", env.JikanApi, id))
	if err != nil {
		return dtos.GetAnimeById{}, err
	}

	if resp.Status == "404" {
		return dtos.GetAnimeById{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.GetAnimeById{}, err
	}

	responseWrapper := models.ResponseWrapper[dtos.GetAnimeById]{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return dtos.GetAnimeById{}, err
	}

	return responseWrapper.Data, nil
}

func GetAnimeStatisticsById(id string) (dtos.GetAnimeStatisticsById, error) {
	resp, err := http.Get(fmt.Sprintf("%sanime/%s/statistics", env.JikanApi, id))
	if err != nil {
		return dtos.GetAnimeStatisticsById{}, err
	}

	if resp.Status == "404" {
		return dtos.GetAnimeStatisticsById{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.GetAnimeStatisticsById{}, err
	}

	responseWrapper := models.ResponseWrapper[dtos.GetAnimeStatisticsById]{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return dtos.GetAnimeStatisticsById{}, err
	}

	return responseWrapper.Data, nil
}
