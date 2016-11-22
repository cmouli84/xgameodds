package domain

import (
	"time"
)

// Event struct
type Event struct {
	ID                int
	GameDate          time.Time
	HomeTeamName      string
	AwayTeamName      string
	HomeOdds          float64
	HomeTeamScore     int
	AwayTeamScore     int
	SonnyMooreRanking struct {
		HomeRanking    float64
		AwayRanking    float64
		SonnyMooreOdds float64
	}
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
