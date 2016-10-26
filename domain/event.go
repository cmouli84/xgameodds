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
	HomeOdds      float32
	HomeTeamScore int
	AwayTeamScore int
}
