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
	parsedStartDate, parseerr := time.ParseInLocation("2006-01-02", date, time.UTC)
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

			if (event.Odd.HomeOdd != "") && !strings.HasPrefix(event.Odd.HomeOdd, "pk") && !strings.HasPrefix(event.Odd.HomeOdd, "N") {
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

			var homeWins, homeLoses, awayWins, awayLoses int
			if strings.Contains(event.Standings.Home.ShortRecord, "-") {
				recs := strings.Split(event.Standings.Home.ShortRecord, "-")
				homeWins, _ = strconv.Atoi(recs[0])
				homeLoses, _ = strconv.Atoi(recs[1])
			}
			if strings.Contains(event.Standings.Away.ShortRecord, "-") {
				recs := strings.Split(event.Standings.Away.ShortRecord, "-")
				awayWins, _ = strconv.Atoi(recs[0])
				awayLoses, _ = strconv.Atoi(recs[1])
			}

			var homeScore = -999999
			var awayScore = -999999
			if eventDate.Before(time.Now()) {
				homeScore = event.BoxScore.Score.Home.Score
				awayScore = event.BoxScore.Score.Away.Score
			}

			domainEvent := domain.Event{
				ID: event.ID,
				HomeTeam: domain.Team{
					Name:    event.HomeTeam.FullName,
					LogoURL: event.HomeTeam.Logos.Tiny,
					Wins:    homeWins,
					Loses:   homeLoses,
				},
				AwayTeam: domain.Team{
					Name:    event.AwayTeam.FullName,
					LogoURL: event.AwayTeam.Logos.Tiny,
					Wins:    awayWins,
					Loses:   awayLoses,
				},
				HomeTeamScore:    homeScore,
				AwayTeamScore:    awayScore,
				WestgateHomeOdds: homeOdds,
				GameDate:         eventDate.Local(),
			}
			filteredEvents = append(filteredEvents, domainEvent)
		}
	}

	return filteredEvents
}
