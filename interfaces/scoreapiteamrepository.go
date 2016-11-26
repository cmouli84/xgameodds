package interfaces

import "github.com/cmouli84/xgameodds/infrastructure"
import "github.com/cmouli84/xgameodds/domain"

// TeamStatsInterface interface
type TeamStatsInterface interface {
	GetNcaabTeamStats() []infrastructure.EventStat
	GetNflTeamStats() []infrastructure.EventStat
}

// ScoreAPITeamRepo struct
type ScoreAPITeamRepo struct {
	teamStatsInterface TeamStatsInterface
}

type getTeamStats func() []infrastructure.EventStat

// NewScoreAPITeamRepo function
func NewScoreAPITeamRepo(teamStatsInterface TeamStatsInterface) *ScoreAPITeamRepo {
	scoreAPITeamRepo := new(ScoreAPITeamRepo)
	scoreAPITeamRepo.teamStatsInterface = teamStatsInterface
	return scoreAPITeamRepo
}

// GetNflTeamStats function
func (scoreAPITeamRepo *ScoreAPITeamRepo) GetNflTeamStats() map[string][]domain.TeamStat {
	return scoreAPITeamRepo.getTeamStats(scoreAPITeamRepo.teamStatsInterface.GetNflTeamStats)
}

// GetNcaabTeamStats function
func (scoreAPITeamRepo *ScoreAPITeamRepo) GetNcaabTeamStats() map[string][]domain.TeamStat {
	return scoreAPITeamRepo.getTeamStats(scoreAPITeamRepo.teamStatsInterface.GetNcaabTeamStats)
}

// getTeamStats function
func (scoreAPITeamRepo *ScoreAPITeamRepo) getTeamStats(getTeamStatsFn getTeamStats) map[string][]domain.TeamStat {
	eventStats := getTeamStatsFn()

	teamStat := make(map[string][]domain.TeamStat)
	for _, eventStat := range eventStats {
		homeTeamStat := domain.TeamStat{}
		homeTeamStat.OpponentTeamName = eventStat.AwayTeamName
		homeTeamStat.HomeTeamScore = eventStat.HomeScore
		homeTeamStat.AwayTeamScore = eventStat.AwayScore
		homeTeamStat.WestgateHomeOdds = eventStat.HomeOdds
		homeTeamStat.EventAtHome = true

		awayTeamStat := domain.TeamStat{}
		awayTeamStat.OpponentTeamName = eventStat.HomeTeamName
		awayTeamStat.HomeTeamScore = eventStat.HomeScore
		awayTeamStat.AwayTeamScore = eventStat.AwayScore
		awayTeamStat.WestgateHomeOdds = eventStat.HomeOdds
		awayTeamStat.EventAtHome = false

		teamStat[eventStat.HomeTeamName] = append(teamStat[eventStat.HomeTeamName], homeTeamStat)
		teamStat[eventStat.AwayTeamName] = append(teamStat[eventStat.AwayTeamName], awayTeamStat)
	}

	return teamStat
}
