package domain

import (
	"time"
)

// Event struct
type Event struct {
	ID            int
	GameDate      time.Time
	HomeTeamName  string
	AwayTeamName  string
	HomeOdds      float64
	HomeTeamScore int
	AwayTeamScore int
}

// EventsInterface interface
type EventsInterface interface {
	GetEventsByDate(eventDate string) []Event
}
