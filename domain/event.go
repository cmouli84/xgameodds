package domain

import (
	"time"
)

// Events type
type Events []Event

// Len function
func (events Events) Len() int {
	return len(events)
}

// Less function
func (events Events) Less(i, j int) bool {
	return events[i].GameDate.Before(events[j].GameDate)
}

// Swap function
func (events Events) Swap(i, j int) {
	events[i], events[j] = events[j], events[i]
}

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
	HomeRecord TeamStats       `json:"homeRecord"`
	AwayRecord TeamStats       `json:"awayRecord"`
	TeamTrends EventTeamTrends `json:"teamTrends"`
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

// TeamStat struct
type TeamStat struct {
	OpponentTeamName string    `json:"opponentTeamName"`
	EventAtHome      bool      `json:"eventAtHome"`
	WestgateHomeOdds float64   `json:"westgateHomeOdds"`
	HomeTeamScore    int       `json:"homeTeamScore"`
	AwayTeamScore    int       `json:"awayTeamScore"`
	EventDate        time.Time `json:"eventDate"`
}

// TeamStats struct
type TeamStats struct {
	Events []TeamStat `json:"events"`
}

// TrendValue struct
type TrendValue struct {
	HomeValue string `json:"homeValue"`
	AwayValue string `json:"awayValue"`
}

// TeamTrendSet struct
type TeamTrendSet struct {
	All                TrendValue `json:"all"`
	Home               TrendValue `json:"home"`
	Away               TrendValue `json:"away"`
	Favorite           TrendValue `json:"favorite"`
	HomeFavorite       TrendValue `json:"homeFavorite"`
	AwayFavorite       TrendValue `json:"awayFavorite"`
	Underdog           TrendValue `json:"underdog"`
	HomeUnderdog       TrendValue `json:"homeUnderdog"`
	AwayUnderdog       TrendValue `json:"awayUnderdog"`
	ConferenceGames    TrendValue `json:"conferenceGames"`
	NonConferenceGames TrendValue `json:"nonConferenceGames"`
	DivisionGames      TrendValue `json:"divisionGames"`
	NonDivisionGames   TrendValue `json:"nonDivisionGames"`
}

// EventTeamTrends struct
type EventTeamTrends struct {
	ATS       TeamTrendSet `json:"ats"`
	OverUnder TeamTrendSet `json:"ou"`
}

// EventsRepository interface
type EventsRepository interface {
	GetNflEventsByDate(eventDate string) []Event
	GetNcaabEventsByDate(eventDate string) []Event
}

// TeamRepository interface
type TeamRepository interface {
	GetNflTeamStats() map[string][]TeamStat
	GetNcaabTeamStats() map[string][]TeamStat
}

// TeamTrendsRepository interface
type TeamTrendsRepository interface {
	GetNflTeamTrends() TeamTrends
	GetNcaabTeamTrends() TeamTrends
}
