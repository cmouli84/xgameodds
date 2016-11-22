package interfaces

import (
	"fmt"
	"time"

	"strconv"
	"strings"

	"github.com/cmouli84/xgameodds/domain"
	"github.com/cmouli84/xgameodds/infrastructure"
)

// ScoreAPIInterface interface
type ScoreAPIInterface interface {
	GetNflSchedule() infrastructure.ScoreAPISchedule
	GetNflEvents(eventIds []int) []infrastructure.ScoreAPIEvent
	GetNcaabSchedule() infrastructure.ScoreAPISchedule
	GetNcaabEvents(eventIds []int) []infrastructure.ScoreAPIEvent
}

type getSchedule func() infrastructure.ScoreAPISchedule

type getEvents func(eventIds []int) []infrastructure.ScoreAPIEvent

// ScoreAPIRepo struct
type ScoreAPIRepo struct {
	scoreAPIInterface ScoreAPIInterface
}

// NewScoreAPIRepo function
func NewScoreAPIRepo(scoreAPIInterface ScoreAPIInterface) *ScoreAPIRepo {
	scoreAPIRepo := new(ScoreAPIRepo)
	scoreAPIRepo.scoreAPIInterface = scoreAPIInterface
	return scoreAPIRepo
}

// GetNcaabEventsByDate function
func (scoreAPIRepo *ScoreAPIRepo) GetNcaabEventsByDate(date string) []domain.Event {
	return scoreAPIRepo.getEventsByDate(date, scoreAPIRepo.scoreAPIInterface.GetNcaabSchedule, scoreAPIRepo.scoreAPIInterface.GetNcaabEvents)
}

// GetNflEventsByDate function
func (scoreAPIRepo *ScoreAPIRepo) GetNflEventsByDate(date string) []domain.Event {
	return scoreAPIRepo.getEventsByDate(date, scoreAPIRepo.scoreAPIInterface.GetNflSchedule, scoreAPIRepo.scoreAPIInterface.GetNflEvents)
}

// getEventsByDate function
func (scoreAPIRepo *ScoreAPIRepo) getEventsByDate(date string, getScheduleFn getSchedule, getEventsFn getEvents) []domain.Event {
	schedule := getScheduleFn()
	parsedStartDate, parseerr := time.ParseInLocation("2006-01-02", date, time.Local)
	if parseerr != nil {
		fmt.Println(parseerr)
		return []domain.Event{}
	}

	parsedEndDate := parsedStartDate.Add(time.Hour * 24)

	eventIds := make([]int, 0)
	for _, season := range schedule.CurrentSeason {
		if (season.StartDate.Before(parsedStartDate) && season.EndDate.After(parsedStartDate)) ||
			(season.StartDate.Before(parsedEndDate) && season.EndDate.After(parsedEndDate)) {
			eventIds = append(eventIds, season.EventIds...)
		}
	}

	events := getEventsFn(eventIds)
	filteredEvents := make([]domain.Event, 0)
	var homeOdds float64
	for _, event := range events {
		eventDate, parseerr := time.Parse(time.RFC1123Z, event.GameDate)
		if parseerr != nil {
			fmt.Println(parseerr)
			continue
		}

		if eventDate.After(parsedStartDate) && eventDate.Before(parsedEndDate) {
			homeOdds = -999999
			var odderr error

			if !strings.HasPrefix(event.Odd.HomeOdd, "pk") && !strings.HasPrefix(event.Odd.HomeOdd, "N") {
				if strings.HasPrefix(event.Odd.HomeOdd, "T") {
					homeOdds, odderr = strconv.ParseFloat(event.Odd.AwayOdd, 64)
					if odderr != nil {
						fmt.Println(odderr)
					}
					homeOdds *= -1
				} else {
					homeOdds, odderr = strconv.ParseFloat(event.Odd.HomeOdd, 64)
					if odderr != nil {
						fmt.Println(odderr)
					}
				}
			}

			domainEvent := domain.Event{
				ID:            event.ID,
				HomeTeamName:  event.HomeTeam.FullName,
				AwayTeamName:  event.AwayTeam.FullName,
				HomeTeamScore: event.BoxScore.Score.Home.Score,
				AwayTeamScore: event.BoxScore.Score.Away.Score,
				HomeOdds:      homeOdds,
				GameDate:      eventDate.Local(),
			}
			filteredEvents = append(filteredEvents, domainEvent)
		}
	}

	return filteredEvents
}
