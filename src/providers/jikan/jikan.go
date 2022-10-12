package jikan

import (
	"encoding/json"
	"fmt"
	"io"
	"jikan-consumer/src/config/env"
	"jikan-consumer/src/models"
	"net/http"
)

func SendRequest[T any](path string) models.ResultModel {
	resp, err := http.Get(fmt.Sprintf("%s%s", env.JikanApi, path))
	if err != nil {
		return models.ResultModel{
			Success:    false,
			StatusCode: 400,
			Errors:     []string{err.Error()},
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

	responseWrapper := models.ResponseWrapper[T]{}

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
