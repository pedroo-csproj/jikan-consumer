package models

type ResponseWrapper[T any] struct {
	Data   T   `json:"data"`
	Status int `json:"status"`
}
