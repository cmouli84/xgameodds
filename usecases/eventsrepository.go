package usecases

import (
	"math"
	"strings"

	"github.com/cmouli84/xgameodds/domain"
)

// EventsInteractor struct
type EventsInteractor struct {
	EventsRepository     domain.EventsRepository
	SonnyMooreRepository domain.SonnyMooreRepository
}

const sonnyMooreHomeAdvantage float64 = 2

// GetNflEventsByDate function
func (interactor *EventsInteractor) GetNflEventsByDate(eventDate string) []domain.Event {
	events := interactor.EventsRepository.GetNflEventsByDate(eventDate)

	sonnyMooreRanking := interactor.SonnyMooreRepository.GetSonnyMooreNflRanking()

	for index := range events {
		awayRanking := sonnyMooreRanking[strings.ToUpper(events[index].AwayTeamName)]
		homeRanking := sonnyMooreRanking[strings.ToUpper(events[index].HomeTeamName)]
		homeOdds := awayRanking - homeRanking - sonnyMooreHomeAdvantage

		events[index].SonnyMooreRanking.AwayRanking = awayRanking
		events[index].SonnyMooreRanking.HomeRanking = homeRanking
		events[index].SonnyMooreRanking.SonnyMooreOdds = Round(homeOdds, 2)
	}

	return events
}

// Round func
func Round(input, places float64) float64 {

	pow := math.Pow(10, places)
	input = input * pow

	if input < 0 {
		return math.Ceil(input-0.5) / pow
	}
	return math.Floor(input+0.5) / pow
}
