package infrastructure

import "time"

// ScoreAPISchedule struct
type ScoreAPISchedule struct {
	CurrentSeason []Group `json:"current_season"`
	CurrentGroup  Group   `json:"current_group"`
}

// Group struct
type Group struct {
	GUID       string    `json:"guid"`
	ID         string    `json:"id"`
	Label      string    `json:"label"`
	SeasonType string    `json:"season_type"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	EventIds   []int     `json:"event_ids"`
}
