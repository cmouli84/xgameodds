package interfaces

import (
	"github.com/patrickmn/go-cache"
)

// CachedTeamNameMapRepo struct
type CachedTeamNameMapRepo struct {
	teamNameMapping TeamnameMappingInterface
	ncaabDataCache  *cache.Cache
}

// NewCachedTeamNameMapRepo function
func NewCachedTeamNameMapRepo(teamNameMapping TeamnameMappingInterface, ncaabDataCache *cache.Cache) *CachedTeamNameMapRepo {
	cachedTeamNameMapRepo := new(CachedTeamNameMapRepo)
	cachedTeamNameMapRepo.teamNameMapping = teamNameMapping
	cachedTeamNameMapRepo.ncaabDataCache = ncaabDataCache
	return cachedTeamNameMapRepo
}

// TeamNameMaps struct
type TeamNameMaps struct {
	SonnyMooreMap map[string]string
	TeamTrendMap  map[string]string
}

// GetNcaabTeamNames function
func (cachedTeamNameMapRepo *CachedTeamNameMapRepo) GetNcaabTeamNames() (map[string]string, map[string]string) {
	var teamNameInterface interface{}
	var teamNameMaps *TeamNameMaps
	var found bool

	teamNameInterface, found = cachedTeamNameMapRepo.ncaabDataCache.Get("TEAMNAMEMAP")

	if found {
		teamNameMaps = teamNameInterface.(*TeamNameMaps)
	} else {
		sonnyTeamMap, teamTrendMap := cachedTeamNameMapRepo.teamNameMapping.GetNcaabTeamNames()
		teamNameMaps = &TeamNameMaps{sonnyTeamMap, teamTrendMap}

		cachedTeamNameMapRepo.ncaabDataCache.Set("TEAMNAMEMAP", teamNameMaps, cache.DefaultExpiration)
	}
	return teamNameMaps.SonnyMooreMap, teamNameMaps.TeamTrendMap
}
