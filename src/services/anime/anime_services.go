package services

import (
	"fmt"
	"jikan-consumer/src/models"
	"jikan-consumer/src/providers/jikan"
	"jikan-consumer/src/services/anime/dtos"
)

func GetAnimeById(id string) models.ResultModel {
	return jikan.SendRequest[dtos.GetAnimeById](fmt.Sprintf("anime/%s/full", id))
}

func GetAnimeStatisticsById(id string) models.ResultModel {
	return jikan.SendRequest[dtos.GetAnimeStatisticsById](fmt.Sprintf("anime/%s/statistics", id))
}
