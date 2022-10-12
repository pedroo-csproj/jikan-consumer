package services

import (
	"fmt"
	"jikan-consumer/src/models"
	"jikan-consumer/src/providers/jikan"
	"jikan-consumer/src/services/club/dtos"
)

func GetClubById(id string) models.ResultModel {
	return jikan.SendRequest[dtos.GetClubById](fmt.Sprintf("clubs/%s", id))
}
