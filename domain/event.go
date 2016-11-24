package domain

import (
	"time"
)

// Event struct
type Event struct {
	ID                int       `json:"id"`
	GameDate          time.Time `json:"gameDate"`
	HomeTeam          Team      `json:"homeTeam"`
	AwayTeam          Team      `json:"awayTeam"`
	WestgateHomeOdds  float64   `json:"westgateHomeOdds"`
	HomeTeamScore     int       `json:"homeTeamScore"`
	AwayTeamScore     int       `json:"awayTeamScore"`
	SonnyMooreRanking struct {
		HomeRanking    float64 `json:"homeRanking"`
		AwayRanking    float64 `json:"awayRanking"`
		SonnyMooreOdds float64 `json:"sonnyMooreOdds"`
	} `json:"sonnyMooreRanking"`
}

// Team struct
type Team struct {
	Name    string `json:"name"`
	LogoURL string `json:"logoUrl"`
	Wins    int    `json:"wins"`
	Loses   int    `json:"loses"`
}

// PersistedRanking struct
type PersistedRanking struct {
	EventID     int
	HomeRanking float64
	AwayRanking float64
}

// EventsRepository interface
type EventsRepository interface {
	GetNflEventsByDate(eventDate string) []Event
	GetNcaabEventsByDate(eventDate string) []Event
}
