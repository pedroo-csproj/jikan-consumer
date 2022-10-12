package models

type ResultModel struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"statusCode"`
	Errors     []string    `json:"errors"`
	Data       interface{} `json:"data"`
}
