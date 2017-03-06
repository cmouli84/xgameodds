package interfaces

import (
	"github.com/patrickmn/go-cache"
	"github.com/cmouli84/xgameodds/domain"
)

// CachedTeamTrendRepo struct
type CachedTeamTrendRepo struct {
	teamTrendInterface TeamTrendsInterface
	nflDataCache     *cache.Cache
	ncaabDataCache   *cache.Cache
}

type getTeamTrendsRepo func() domain.TeamTrends

// NewCachedTeamTrendRepo function
func NewCachedTeamTrendRepo(teamTrendInterface TeamTrendsInterface, nflDataCache *cache.Cache, ncaabDataCache *cache.Cache) *CachedTeamTrendRepo {
	cachedTeamTrendRepo := new(CachedTeamTrendRepo)
	cachedTeamTrendRepo.teamTrendInterface = teamTrendInterface
	cachedTeamTrendRepo.nflDataCache = nflDataCache
	cachedTeamTrendRepo.ncaabDataCache = ncaabDataCache
	return cachedTeamTrendRepo
}

// GetNflTeamTrends function
func (cachedTeamTrendRepo *CachedTeamTrendRepo) GetNflTeamTrends() domain.TeamTrends {
	return cachedTeamTrendRepo.getTeamTrends(cachedTeamTrendRepo.teamTrendInterface.GetNflTeamTrends, cachedTeamTrendRepo.nflDataCache)
}

// GetNcaabTeamTrends function
func (cachedTeamTrendRepo *CachedTeamTrendRepo) GetNcaabTeamTrends() domain.TeamTrends {
	return cachedTeamTrendRepo.getTeamTrends(cachedTeamTrendRepo.teamTrendInterface.GetNcaabTeamTrends, cachedTeamTrendRepo.ncaabDataCache)
}

// getTeamTrends function
func (cachedTeamTrendRepo *CachedTeamTrendRepo) getTeamTrends(getTeamTrendsRepoFn getTeamTrendsRepo, dataCache *cache.Cache) domain.TeamTrends {
	var teamTrendsInterface interface{}
	var teamTrends *domain.TeamTrends
	var found bool

	teamTrendsInterface, found = dataCache.Get("TEAMTRENDS")

	if found {
		teamTrends = teamTrendsInterface.(*domain.TeamTrends)
	} else {
		teamTrendsRepo := getTeamTrendsRepoFn()
		teamTrends = &teamTrendsRepo

		dataCache.Set("TEAMTRENDS", teamTrends, cache.DefaultExpiration)
	}
	return *teamTrends
}
