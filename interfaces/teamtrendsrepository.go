package interfaces

import "github.com/cmouli84/xgameodds/domain"

// TeamTrendsInterface interface
type TeamTrendsInterface interface {
	GetNflTeamTrends() domain.TeamTrends
	GetNcaabTeamTrends() domain.TeamTrends
}

// TeamTrendsRepo struct
type TeamTrendsRepo struct {
	teamTrendsInterface      TeamTrendsInterface
	teamnameMappingInterface TeamnameMappingInterface
}

// NewTeamTrendsRepo function
func NewTeamTrendsRepo(teamTrendsInterface TeamTrendsInterface, teamnameMappingInterface TeamnameMappingInterface) *TeamTrendsRepo {
	teamTrendsRepo := new(TeamTrendsRepo)
	teamTrendsRepo.teamTrendsInterface = teamTrendsInterface
	teamTrendsRepo.teamnameMappingInterface = teamnameMappingInterface
	return teamTrendsRepo
}

// GetNflTeamTrends function
func (teamTrendsRepo *TeamTrendsRepo) GetNflTeamTrends() domain.TeamTrends {
	return teamTrendsRepo.teamTrendsInterface.GetNflTeamTrends()
}

// GetNcaabTeamTrends function
func (teamTrendsRepo *TeamTrendsRepo) GetNcaabTeamTrends() domain.TeamTrends {
	teamTrendMap := teamTrendsRepo.teamTrendsInterface.GetNcaabTeamTrends()
	_, teamMap := teamTrendsRepo.teamnameMappingInterface.GetNcaabTeamNames()

	result := domain.TeamTrends(make(map[string]domain.Trend))
	for k, v := range teamTrendMap {
		result[teamMap[k]] = v
	}

	return result
}
