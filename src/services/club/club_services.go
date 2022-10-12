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

func GetClubById(id string) models.ResultModel {
	resp, err := http.Get(fmt.Sprintf("%sclubs/%s", env.JikanApi, id))
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
			Errors:     []string{"club doens't exists"},
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

	responseWrapper := models.ResponseWrapper[dtos.GetClubById]{}

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
