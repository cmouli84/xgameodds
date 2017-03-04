package interfaces

import "github.com/cmouli84/xgameodds/infrastructure"

// TeamTrendsInterface interface
type TeamTrendsInterface interface {
	GetNflTeamTrends() infrastructure.TeamTrends
	GetNcaabTeamTrends() infrastructure.TeamTrends
}

// TeamTrendsRepo struct
type TeamTrendsRepo struct {
	teamTrendsInterface      TeamTrendsInterface
	teamnameMappingInterface TeamnameMappingInterface
}

// NewTeamTrendsInterface function
func NewTeamTrendsInterface(teamTrendsInterface TeamTrendsInterface, teamnameMappingInterface TeamnameMappingInterface) *TeamTrendsRepo {
	teamTrendsRepo := new(TeamTrendsRepo)
	teamTrendsRepo.teamTrendsInterface = teamTrendsInterface
	teamTrendsRepo.teamnameMappingInterface = teamnameMappingInterface
	return teamTrendsRepo
}

// GetNflTeamTrends function
func (teamTrendsRepo *TeamTrendsRepo) GetNflTeamTrends() infrastructure.TeamTrends {
	return teamTrendsRepo.teamTrendsInterface.GetNflTeamTrends()
}

// GetNcaabTeamTrends function
func (teamTrendsRepo *TeamTrendsRepo) GetNcaabTeamTrends() infrastructure.TeamTrends {
	teamTrendMap := teamTrendsRepo.teamTrendsInterface.GetNcaabTeamTrends()
	_, teamMap := teamTrendsRepo.teamnameMappingInterface.GetNcaabTeamNames()

	result := infrastructure.TeamTrends(make(map[string]infrastructure.Trend))
	for k, v := range teamTrendMap {
		result[teamMap[k]] = v
	}

	return result
}
