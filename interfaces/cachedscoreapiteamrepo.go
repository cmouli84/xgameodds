package interfaces

import (
	"github.com/cmouli84/xgameodds/domain"
	"github.com/patrickmn/go-cache"
)

// CachedScoreAPITeamRepo struct
type CachedScoreAPITeamRepo struct {
	scoreAPITeamRepo domain.TeamRepository
	nflDataCache     *cache.Cache
	ncaabDataCache   *cache.Cache
}

type getTeamStatsRepo func() map[string][]domain.TeamStat

// NewCachedScoreAPITeamRepo function
func NewCachedScoreAPITeamRepo(scoreAPITeamRepo domain.TeamRepository, nflDataCache *cache.Cache, ncaabDataCache *cache.Cache) *CachedScoreAPITeamRepo {
	cachedScoreAPITeamRepo := new(CachedScoreAPITeamRepo)
	cachedScoreAPITeamRepo.scoreAPITeamRepo = scoreAPITeamRepo
	cachedScoreAPITeamRepo.nflDataCache = nflDataCache
	cachedScoreAPITeamRepo.ncaabDataCache = ncaabDataCache
	return cachedScoreAPITeamRepo
}

// GetNflTeamStats function
func (cachedScoreAPITeamRepo *CachedScoreAPITeamRepo) GetNflTeamStats() map[string][]domain.TeamStat {
	return cachedScoreAPITeamRepo.getTeamStats(cachedScoreAPITeamRepo.scoreAPITeamRepo.GetNflTeamStats, cachedScoreAPITeamRepo.nflDataCache)
}

// GetNcaabTeamStats function
func (cachedScoreAPITeamRepo *CachedScoreAPITeamRepo) GetNcaabTeamStats() map[string][]domain.TeamStat {
	return cachedScoreAPITeamRepo.getTeamStats(cachedScoreAPITeamRepo.scoreAPITeamRepo.GetNcaabTeamStats, cachedScoreAPITeamRepo.ncaabDataCache)
}

// getTeamStats function
func (cachedScoreAPITeamRepo *CachedScoreAPITeamRepo) getTeamStats(getTeamStatsFn getTeamStatsRepo, dataCache *cache.Cache) map[string][]domain.TeamStat {
	var teamStatsInterface interface{}
	var teamStats *map[string][]domain.TeamStat
	var found bool

	teamStatsInterface, found = dataCache.Get("TEAMSTAT")

	if found {
		teamStats = teamStatsInterface.(*map[string][]domain.TeamStat)
	} else {
		teamStatsRepo := getTeamStatsFn()
		teamStats = &teamStatsRepo

		dataCache.Set("TEAMSTAT", teamStats, cache.DefaultExpiration)
	}
	return *teamStats
}
