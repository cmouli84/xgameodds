package interfaces

import (
	"time"

	"github.com/cmouli84/xgameodds/domain"
	"github.com/cmouli84/xgameodds/infrastructure"
)

// ScoreAPIInterface interface
type ScoreAPIInterface interface {
	GetSchedule() infrastructure.ScoreAPISchedule
	GetEvents(eventIds []int) []infrastructure.ScoreAPIEvent
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
	schedule := scoreAPIRepo.scoreAPIInterface.GetSchedule()
	parsedStartDate, _ := time.ParseInLocation("2016-01-02", date, time.Local)
	parsedEndDate := parsedStartDate.Add(time.Hour * 24)

	eventIds := make([]int, 0)
	for _, season := range schedule.CurrentSeason {
		if (season.StartDate.Before(parsedStartDate) && season.EndDate.After(parsedStartDate)) ||
			(season.StartDate.Before(parsedEndDate) && season.EndDate.After(parsedEndDate)) {
			eventIds = append(eventIds, season.EventIds...)
		}
	}

	events := scoreAPIRepo.scoreAPIInterface.GetEvents(eventIds)
	filteredEvents := make([]domain.Event, 0)
	for _, event := range events {
		eventDate, _ := time.Parse(time.RFC1123Z, event.GameDate)
		if eventDate.After(parsedStartDate) && eventDate.Before(parsedEndDate) {
			domainEvent := domain.Event{ID: event.ID, HomeTeamName: event.HomeTeam.FullName, AwayTeamName: event.AwayTeam.FullName, HomeTeamScore: event.BoxScore.Score.Home.Score, AwayTeamScore: event.BoxScore.Score.Away.Score, HomeOdds: 0 /*event.Odd.HomeOdd*/, GameDate: eventDate /*event.GameDate */}
			filteredEvents = append(filteredEvents, domainEvent)
		}
	}

	return filteredEvents
}
