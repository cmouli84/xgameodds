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
	TeamRepository       domain.TeamRepository
	TeamTrendsRepository domain.TeamTrendsRepository
	SonnyMooreRepository domain.SonnyMooreRepository
	DynamoDbRepository   domain.DynamoDbRepository
}

type getEventsByDate func(eventDate string) []domain.Event

type getTeamStats func() map[string][]domain.TeamStat

type getSonnyMooreRanking func() map[string]float64

type getPersistedRanking func(eventIds []int) map[int]domain.PersistedRanking

type getTeamTrends func() domain.TeamTrends

const sonnyMooreNflHomeAdvantage float64 = 3.2

const sonnyMooreNcaabHomeAdvantage float64 = 0 //3.25

// GetNflEventsByDate function
func (interactor *EventsInteractor) GetNflEventsByDate(eventDate string) []domain.Event {
	return interactor.getEventsByDate(eventDate, interactor.EventsRepository.GetNflEventsByDate, interactor.TeamRepository.GetNflTeamStats, interactor.TeamTrendsRepository.GetNflTeamTrends, interactor.SonnyMooreRepository.GetSonnyMooreNflRanking, interactor.DynamoDbRepository.GetNflPersistedRanking, sonnyMooreNflHomeAdvantage)
}

// GetNcaabEventsByDate function
func (interactor *EventsInteractor) GetNcaabEventsByDate(eventDate string) []domain.Event {
	return interactor.getEventsByDate(eventDate, interactor.EventsRepository.GetNcaabEventsByDate, interactor.TeamRepository.GetNcaabTeamStats, interactor.TeamTrendsRepository.GetNcaabTeamTrends, interactor.SonnyMooreRepository.GetSonnyMooreNcaabRanking, interactor.DynamoDbRepository.GetNcaabPersistedRanking, sonnyMooreNcaabHomeAdvantage)
}

// getEventsByDate function
func (interactor *EventsInteractor) getEventsByDate(eventDate string, getEventByDateFn getEventsByDate, getTeamStatsFn getTeamStats, getTeamTrendsFn getTeamTrends, getSonnyMooreRankingFn getSonnyMooreRanking, getPersistedRankingFn getPersistedRanking, homeAdvantage float64) []domain.Event {
	events := getEventByDateFn(eventDate)
	teamStats := make(map[string][]domain.TeamStat)
	teamTrends := domain.TeamTrends{}
	if len(events) > 0 {
		teamStats = getTeamStatsFn()
		teamTrends = getTeamTrendsFn()
	}

	sonnyMooreRanking := getSonnyMooreRankingFn()

	pastEvents := getPastEvents(events)
	pastRanking := getPersistedRankingFn(pastEvents)
	currentTime := time.Now()

	for index, event := range events {
		var awayRanking float64 = -999999
		var homeRanking float64 = -999999
		if event.GameDate.Before(currentTime) {
			if pastRanking[event.ID].HomeRanking != 0 && pastRanking[event.ID].AwayRanking != 0 {
				awayRanking = pastRanking[event.ID].AwayRanking
				homeRanking = pastRanking[event.ID].HomeRanking
			}
		} else {
			awayRanking = sonnyMooreRanking[strings.ToUpper(events[index].AwayTeam.Name)]
			homeRanking = sonnyMooreRanking[strings.ToUpper(events[index].HomeTeam.Name)]
		}

		if (event.GameDate.Year() > currentTime.Year()) || ((event.GameDate.Year() == currentTime.Year()) && (event.GameDate.YearDay() >= currentTime.YearDay())) {
			events[index].HomeRecord.Events = teamStats[event.HomeTeam.Name]
			events[index].AwayRecord.Events = teamStats[event.AwayTeam.Name]
			events[index].TeamTrends = getTeamTrendsByTeam(teamTrends, event.HomeTeam.Name, event.AwayTeam.Name)
		}

		var homeOdds float64 = -999999
		if awayRanking != -999999 && homeRanking != -999999 {
			homeOdds = awayRanking - homeRanking - homeAdvantage
		}

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

func getTeamTrendsByTeam(teamTrends domain.TeamTrends, homeTeamName, awayTeamName string) domain.EventTeamTrends {
	eventTeamTrends := domain.EventTeamTrends{}

	if homeTeamTrend, ok := teamTrends[homeTeamName]; ok {
		eventTeamTrends.ATS.All.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.All)
		eventTeamTrends.ATS.Away.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.Away)
		eventTeamTrends.ATS.AwayFavorite.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.AwayFavorite)
		eventTeamTrends.ATS.AwayUnderdog.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.AwayUnderdog)
		eventTeamTrends.ATS.ConferenceGames.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.ConferenceGames)
		eventTeamTrends.ATS.DivisionGames.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.DivisionGames)
		eventTeamTrends.ATS.Favorite.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.Favorite)
		eventTeamTrends.ATS.Home.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.Home)
		eventTeamTrends.ATS.HomeFavorite.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.HomeFavorite)
		eventTeamTrends.ATS.HomeUnderdog.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.HomeUnderdog)
		eventTeamTrends.ATS.NonConferenceGames.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.NonConferenceGames)
		eventTeamTrends.ATS.NonDivisionGames.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.NonDivisionGames)
		eventTeamTrends.ATS.Underdog.HomeValue = getTeamTrendValue(homeTeamTrend.Ats.Underdog)

		eventTeamTrends.OverUnder.All.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.All)
		eventTeamTrends.OverUnder.Away.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.Away)
		eventTeamTrends.OverUnder.AwayFavorite.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.AwayFavorite)
		eventTeamTrends.OverUnder.AwayUnderdog.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.AwayUnderdog)
		eventTeamTrends.OverUnder.ConferenceGames.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.ConferenceGames)
		eventTeamTrends.OverUnder.DivisionGames.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.DivisionGames)
		eventTeamTrends.OverUnder.Favorite.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.Favorite)
		eventTeamTrends.OverUnder.Home.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.Home)
		eventTeamTrends.OverUnder.HomeFavorite.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.HomeFavorite)
		eventTeamTrends.OverUnder.HomeUnderdog.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.HomeUnderdog)
		eventTeamTrends.OverUnder.NonConferenceGames.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.NonConferenceGames)
		eventTeamTrends.OverUnder.NonDivisionGames.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.NonDivisionGames)
		eventTeamTrends.OverUnder.Underdog.HomeValue = getTeamTrendValue(homeTeamTrend.OverUnder.Underdog)
	} else {
		eventTeamTrends.ATS.All.HomeValue = "-"
		eventTeamTrends.ATS.Away.HomeValue = "-"
		eventTeamTrends.ATS.AwayFavorite.HomeValue = "-"
		eventTeamTrends.ATS.AwayUnderdog.HomeValue = "-"
		eventTeamTrends.ATS.ConferenceGames.HomeValue = "-"
		eventTeamTrends.ATS.DivisionGames.HomeValue = "-"
		eventTeamTrends.ATS.Favorite.HomeValue = "-"
		eventTeamTrends.ATS.Home.HomeValue = "-"
		eventTeamTrends.ATS.HomeFavorite.HomeValue = "-"
		eventTeamTrends.ATS.HomeUnderdog.HomeValue = "-"
		eventTeamTrends.ATS.NonConferenceGames.HomeValue = "-"
		eventTeamTrends.ATS.NonDivisionGames.HomeValue = "-"
		eventTeamTrends.ATS.Underdog.HomeValue = "-"

		eventTeamTrends.OverUnder.All.HomeValue = "-"
		eventTeamTrends.OverUnder.Away.HomeValue = "-"
		eventTeamTrends.OverUnder.AwayFavorite.HomeValue = "-"
		eventTeamTrends.OverUnder.AwayUnderdog.HomeValue = "-"
		eventTeamTrends.OverUnder.ConferenceGames.HomeValue = "-"
		eventTeamTrends.OverUnder.DivisionGames.HomeValue = "-"
		eventTeamTrends.OverUnder.Favorite.HomeValue = "-"
		eventTeamTrends.OverUnder.Home.HomeValue = "-"
		eventTeamTrends.OverUnder.HomeFavorite.HomeValue = "-"
		eventTeamTrends.OverUnder.HomeUnderdog.HomeValue = "-"
		eventTeamTrends.OverUnder.NonConferenceGames.HomeValue = "-"
		eventTeamTrends.OverUnder.NonDivisionGames.HomeValue = "-"
		eventTeamTrends.OverUnder.Underdog.HomeValue = "-"
	}

	if awayTeamTrend, ok := teamTrends[awayTeamName]; ok {
		eventTeamTrends.ATS.All.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.All)
		eventTeamTrends.ATS.Away.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.Away)
		eventTeamTrends.ATS.AwayFavorite.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.AwayFavorite)
		eventTeamTrends.ATS.AwayUnderdog.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.AwayUnderdog)
		eventTeamTrends.ATS.ConferenceGames.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.ConferenceGames)
		eventTeamTrends.ATS.DivisionGames.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.DivisionGames)
		eventTeamTrends.ATS.Favorite.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.Favorite)
		eventTeamTrends.ATS.Home.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.Home)
		eventTeamTrends.ATS.HomeFavorite.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.HomeFavorite)
		eventTeamTrends.ATS.HomeUnderdog.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.HomeUnderdog)
		eventTeamTrends.ATS.NonConferenceGames.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.NonConferenceGames)
		eventTeamTrends.ATS.NonDivisionGames.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.NonDivisionGames)
		eventTeamTrends.ATS.Underdog.AwayValue = getTeamTrendValue(awayTeamTrend.Ats.Underdog)

		eventTeamTrends.OverUnder.All.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.All)
		eventTeamTrends.OverUnder.Away.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.Away)
		eventTeamTrends.OverUnder.AwayFavorite.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.AwayFavorite)
		eventTeamTrends.OverUnder.AwayUnderdog.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.AwayUnderdog)
		eventTeamTrends.OverUnder.ConferenceGames.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.ConferenceGames)
		eventTeamTrends.OverUnder.DivisionGames.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.DivisionGames)
		eventTeamTrends.OverUnder.Favorite.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.Favorite)
		eventTeamTrends.OverUnder.Home.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.Home)
		eventTeamTrends.OverUnder.HomeFavorite.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.HomeFavorite)
		eventTeamTrends.OverUnder.HomeUnderdog.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.HomeUnderdog)
		eventTeamTrends.OverUnder.NonConferenceGames.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.NonConferenceGames)
		eventTeamTrends.OverUnder.NonDivisionGames.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.NonDivisionGames)
		eventTeamTrends.OverUnder.Underdog.AwayValue = getTeamTrendValue(awayTeamTrend.OverUnder.Underdog)
	} else {
		eventTeamTrends.ATS.All.AwayValue = "-"
		eventTeamTrends.ATS.Away.AwayValue = "-"
		eventTeamTrends.ATS.AwayFavorite.AwayValue = "-"
		eventTeamTrends.ATS.AwayUnderdog.AwayValue = "-"
		eventTeamTrends.ATS.ConferenceGames.AwayValue = "-"
		eventTeamTrends.ATS.DivisionGames.AwayValue = "-"
		eventTeamTrends.ATS.Favorite.AwayValue = "-"
		eventTeamTrends.ATS.Home.AwayValue = "-"
		eventTeamTrends.ATS.HomeFavorite.AwayValue = "-"
		eventTeamTrends.ATS.HomeUnderdog.AwayValue = "-"
		eventTeamTrends.ATS.NonConferenceGames.AwayValue = "-"
		eventTeamTrends.ATS.NonDivisionGames.AwayValue = "-"
		eventTeamTrends.ATS.Underdog.AwayValue = "-"

		eventTeamTrends.OverUnder.All.AwayValue = "-"
		eventTeamTrends.OverUnder.Away.AwayValue = "-"
		eventTeamTrends.OverUnder.AwayFavorite.AwayValue = "-"
		eventTeamTrends.OverUnder.AwayUnderdog.AwayValue = "-"
		eventTeamTrends.OverUnder.ConferenceGames.AwayValue = "-"
		eventTeamTrends.OverUnder.DivisionGames.AwayValue = "-"
		eventTeamTrends.OverUnder.Favorite.AwayValue = "-"
		eventTeamTrends.OverUnder.Home.AwayValue = "-"
		eventTeamTrends.OverUnder.HomeFavorite.AwayValue = "-"
		eventTeamTrends.OverUnder.HomeUnderdog.AwayValue = "-"
		eventTeamTrends.OverUnder.NonConferenceGames.AwayValue = "-"
		eventTeamTrends.OverUnder.NonDivisionGames.AwayValue = "-"
		eventTeamTrends.OverUnder.Underdog.AwayValue = "-"
	}

	return eventTeamTrends
}

func getTeamTrendValue(trendValue string) string {
	if trendValue == "" {
		return "-"
	}
	return trendValue
}
