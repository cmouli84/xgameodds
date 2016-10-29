package interfaces

import (
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
}

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

// GetEventsByDate function
func (scoreAPIRepo *ScoreAPIRepo) GetEventsByDate(date string) []domain.Event {
	schedule := scoreAPIRepo.scoreAPIInterface.GetNflSchedule()
	parsedStartDate, _ := time.ParseInLocation("2006-01-02", date, time.Local)
	parsedEndDate := parsedStartDate.Add(time.Hour * 24)

	eventIds := make([]int, 0)
	for _, season := range schedule.CurrentSeason {
		if (season.StartDate.Before(parsedStartDate) && season.EndDate.After(parsedStartDate)) ||
			(season.StartDate.Before(parsedEndDate) && season.EndDate.After(parsedEndDate)) {
			eventIds = append(eventIds, season.EventIds...)
		}
	}

	events := scoreAPIRepo.scoreAPIInterface.GetNflEvents(eventIds)
	filteredEvents := make([]domain.Event, 0)
	var homeOdds float64
	for _, event := range events {
		eventDate, _ := time.Parse(time.RFC1123Z, event.GameDate)

		if eventDate.After(parsedStartDate) && eventDate.Before(parsedEndDate) {
			homeOdds = -999999

			if !strings.HasPrefix(event.Odd.HomeOdd, "pk") && !strings.HasPrefix(event.Odd.HomeOdd, "N") {
				if strings.HasPrefix(event.Odd.HomeOdd, "T") {
					homeOdds, _ = strconv.ParseFloat(event.Odd.AwayOdd, 64)
					homeOdds *= -1
				} else {
					homeOdds, _ = strconv.ParseFloat(event.Odd.HomeOdd, 64)
				}
			}

			gameDate, _ := time.Parse(time.RFC1123Z, event.GameDate)

			domainEvent := domain.Event{
				ID:            event.ID,
				HomeTeamName:  event.HomeTeam.FullName,
				AwayTeamName:  event.AwayTeam.FullName,
				HomeTeamScore: event.BoxScore.Score.Home.Score,
				AwayTeamScore: event.BoxScore.Score.Away.Score,
				HomeOdds:      homeOdds,
				GameDate:      gameDate.Local(),
			}
			filteredEvents = append(filteredEvents, domainEvent)
		}
	}

	return filteredEvents
}
