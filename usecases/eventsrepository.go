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

type getEventsByDate func(eventDate string) []domain.Event

type getSonnyMooreRanking func() map[string]float64

type getPersistedRanking func(eventIds []int) map[int]domain.PersistedRanking

const sonnyMooreHomeAdvantage float64 = 2

// GetNflEventsByDate function
func (interactor *EventsInteractor) GetNflEventsByDate(eventDate string) []domain.Event {
	return interactor.getEventsByDate(eventDate, interactor.EventsRepository.GetNflEventsByDate, interactor.SonnyMooreRepository.GetSonnyMooreNflRanking, interactor.DynamoDbRepository.GetNflPersistedRanking)
}

// GetNcaabEventsByDate function
func (interactor *EventsInteractor) GetNcaabEventsByDate(eventDate string) []domain.Event {
	return interactor.getEventsByDate(eventDate, interactor.EventsRepository.GetNcaabEventsByDate, interactor.SonnyMooreRepository.GetSonnyMooreNcaabRanking, interactor.DynamoDbRepository.GetNcaabPersistedRanking)
}

// getEventsByDate function
func (interactor *EventsInteractor) getEventsByDate(eventDate string, getEventByDateFn getEventsByDate, getSonnyMooreRankingFn getSonnyMooreRanking, getPersistedRankingFn getPersistedRanking) []domain.Event {
	events := getEventByDateFn(eventDate)

	sonnyMooreRanking := getSonnyMooreRankingFn()

	pastEvents := getPastEvents(events)
	pastRanking := getPersistedRankingFn(pastEvents)
	currentTime := time.Now()

	var awayRanking, homeRanking float64
	for index, event := range events {
		if event.GameDate.Before(currentTime) {
			awayRanking = pastRanking[event.ID].AwayRanking
			homeRanking = pastRanking[event.ID].HomeRanking
		} else {
			awayRanking = sonnyMooreRanking[strings.ToUpper(events[index].AwayTeam.Name)]
			homeRanking = sonnyMooreRanking[strings.ToUpper(events[index].HomeTeam.Name)]
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
