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

func GetAnimeById(id string) models.ResultModel {
	resp, err := http.Get(fmt.Sprintf("%sanime/%s/full", env.JikanApi, id))
	if err != nil {
		return models.ResultModel{
			Success:    false,
			StatusCode: 400,
			Errors:     []string{err.Error()},
		}
	}

	if resp.Status == "404" {
		return models.ResultModel{
			Success:    false,
			StatusCode: 404,
			Errors:     []string{"anime doens't exists"},
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ResultModel{
			Success:    false,
			StatusCode: 400,
			Errors:     []string{err.Error()},
		}
	}

	responseWrapper := models.ResponseWrapper[dtos.GetAnimeById]{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return models.ResultModel{
			Success:    false,
			StatusCode: 400,
			Errors:     []string{err.Error()},
		}
	}

	return models.ResultModel{
		Success:    true,
		StatusCode: 200,
		Data:       responseWrapper.Data,
	}
}

func GetAnimeStatisticsById(id string) models.ResultModel {
	resp, err := http.Get(fmt.Sprintf("%sanime/%s/statistics", env.JikanApi, id))
	if err != nil {
		return models.ResultModel{
			Success:    false,
			StatusCode: 400,
			Errors:     []string{err.Error()},
		}
	}

	if resp.Status == "404" {
		return models.ResultModel{
			Success:    false,
			StatusCode: 404,
			Errors:     []string{"anime doens't exists"},
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.ResultModel{
			Success:    false,
			StatusCode: 400,
			Errors:     []string{err.Error()},
		}
	}

	responseWrapper := models.ResponseWrapper[dtos.GetAnimeStatisticsById]{}

	err = json.Unmarshal(body, &responseWrapper)
	if err != nil {
		return models.ResultModel{
			Success:    false,
			StatusCode: 400,
			Errors:     []string{err.Error()},
		}
	}

	return models.ResultModel{
		Success:    true,
		StatusCode: 200,
		Data:       responseWrapper.Data,
	}
}
