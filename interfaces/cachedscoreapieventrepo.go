package interfaces

import (
	"github.com/cmouli84/xgameodds/domain"
	"github.com/patrickmn/go-cache"
)

// CachedScoreAPIRepo struct
type CachedScoreAPIRepo struct {
	scoreAPIRepo   domain.EventsRepository
	nflDataCache   *cache.Cache
	ncaabDataCache *cache.Cache
}

type getEventsByDate func(date string) []domain.Event

// NewCachedScoreAPIRepo function
func NewCachedScoreAPIRepo(scoreRepoInterface domain.EventsRepository, nflDataCache *cache.Cache, ncaabDataCache *cache.Cache) *CachedScoreAPIRepo {
	cachedScoreAPIRepo := new(CachedScoreAPIRepo)
	cachedScoreAPIRepo.scoreAPIRepo = scoreRepoInterface
	cachedScoreAPIRepo.nflDataCache = nflDataCache
	cachedScoreAPIRepo.ncaabDataCache = ncaabDataCache
	return cachedScoreAPIRepo
}

// GetNflEventsByDate function
func (cachedScoreAPIRepo *CachedScoreAPIRepo) GetNflEventsByDate(date string) []domain.Event {
	return cachedScoreAPIRepo.getEventsByDate(date, cachedScoreAPIRepo.scoreAPIRepo.GetNflEventsByDate, cachedScoreAPIRepo.nflDataCache)
}

// GetNcaabEventsByDate function
func (cachedScoreAPIRepo *CachedScoreAPIRepo) GetNcaabEventsByDate(date string) []domain.Event {
	return cachedScoreAPIRepo.getEventsByDate(date, cachedScoreAPIRepo.scoreAPIRepo.GetNcaabEventsByDate, cachedScoreAPIRepo.ncaabDataCache)
}

// getEventsByDate function
func (cachedScoreAPIRepo *CachedScoreAPIRepo) getEventsByDate(date string, getEventsByDateFn getEventsByDate, dataCache *cache.Cache) []domain.Event {
	var eventsInterface interface{}
	var events *[]domain.Event
	var found bool

	eventsInterface, found = dataCache.Get(date)

	if found {
		events = eventsInterface.(*[]domain.Event)
	} else {
		eventsRepo := getEventsByDateFn(date)
		events = &eventsRepo

		dataCache.Set(date, events, cache.DefaultExpiration)
	}
	return *events
}
