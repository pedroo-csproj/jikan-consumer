package dtos

type GetClubById struct {
	ID       int    `json:"mal_id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	Category string `json:"category"`
}
