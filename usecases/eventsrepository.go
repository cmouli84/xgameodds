package usecases

import (
	"math"
	"strings"

	"time"

	"github.com/cmouli84/xgameodds/domain"
)

// EventsInteractor struct
type EventsInteractor struct {
	EventsRepository     domain.EventsRepository
	SonnyMooreRepository domain.SonnyMooreRepository
	DynamoDbRepository   domain.DynamoDbRepository
}

const sonnyMooreHomeAdvantage float64 = 2

// GetNflEventsByDate function
func (interactor *EventsInteractor) GetNflEventsByDate(eventDate string) []domain.Event {
	events := interactor.EventsRepository.GetNflEventsByDate(eventDate)

	sonnyMooreRanking := interactor.SonnyMooreRepository.GetSonnyMooreNflRanking()

	pastEvents := getPastEvents(events)
	pastRanking := interactor.DynamoDbRepository.GetNflPersistedRanking(pastEvents)
	currentTime := time.Now()

	var awayRanking, homeRanking float64
	for index, event := range events {
		if event.GameDate.Before(currentTime) {
			awayRanking = pastRanking[event.ID].AwayRanking
			homeRanking = pastRanking[event.ID].HomeRanking
		} else {
			awayRanking = sonnyMooreRanking[strings.ToUpper(events[index].AwayTeamName)]
			homeRanking = sonnyMooreRanking[strings.ToUpper(events[index].HomeTeamName)]
		}

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

func getPastEvents(events []domain.Event) []int {
	currentTime := time.Now()

	eventIds := make([]int, 0)
	for _, event := range events {
		if event.GameDate.Before(currentTime) {
			eventIds = append(eventIds, event.ID)
		}
	}

	return eventIds
}
