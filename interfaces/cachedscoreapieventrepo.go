package interfaces

import (
	"github.com/cmouli84/xgameodds/domain"
	"github.com/patrickmn/go-cache"
)

// CachedScoreAPIRepo struct
type CachedScoreAPIRepo struct {
	scoreAPIRepo domain.EventsRepository
	dataCache    *cache.Cache
}

// NewCachedScoreAPIRepo function
func NewCachedScoreAPIRepo(scoreRepoInterface domain.EventsRepository, datacache *cache.Cache) *CachedScoreAPIRepo {
	cachedScoreAPIRepo := new(CachedScoreAPIRepo)
	cachedScoreAPIRepo.scoreAPIRepo = scoreRepoInterface
    cachedScoreAPIRepo.dataCache = datacache
	return cachedScoreAPIRepo
}

// GetNflEventsByDate function
func (cachedScoreAPIRepo *CachedScoreAPIRepo) GetNflEventsByDate(date string) []domain.Event {
	var nflEventsInterface interface{}
	var nflEvents *[]domain.Event
	var found bool

	nflEventsInterface, found = cachedScoreAPIRepo.dataCache.Get(date)

	if found {
		nflEvents = nflEventsInterface.(*[]domain.Event)
	} else {
		nflEventsRepo := cachedScoreAPIRepo.scoreAPIRepo.GetNflEventsByDate(date)
		nflEvents = &nflEventsRepo

		cachedScoreAPIRepo.dataCache.Set(date, nflEvents, cache.DefaultExpiration)
	}
	return *nflEvents
}
