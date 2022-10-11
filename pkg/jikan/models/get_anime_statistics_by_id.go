package models

type GetAnimeStatisticsById struct {
	Watching    int `json:"watching"`
	Completed   int `json:"completed"`
	OnHold      int `json:"on_hold"`
	Dropped     int `json:"dropped"`
	PlanToWatch int `json:"plan_to_watch"`
}
