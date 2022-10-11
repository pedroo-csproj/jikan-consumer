package models

type GetAnimeById struct {
	ID            int     `json:"mal_id"`
	Title         string  `json:"title"`
	EnglishTitle  string  `json:"title_english"`
	JapaneseTitle string  `json:"title_japanese"`
	Type          string  `json:"type"`
	Status        string  `json:"status"`
	Airing        bool    `json:"airing"`
	Rating        string  `json:"rating"`
	Score         float32 `json:"score"`
	Rank          int     `json:"rank"`
	Members       int     `json:"members"`
	Episodes      int     `json:"episodes"`
	Url           string  `json:"url"`
}
