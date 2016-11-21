package interfaces

import (
	"strconv"

	"github.com/cmouli84/xgameodds/domain"
	"github.com/patrickmn/go-cache"
)

// CachedPersistedRankingRepo struct
type CachedPersistedRankingRepo struct {
	dynamodbRepo domain.DynamoDbRepository
	dataCache    *cache.Cache
}

// NewCachedPersistedRankingRepo function
func NewCachedPersistedRankingRepo(dynamodbRepo domain.DynamoDbRepository, datacache *cache.Cache) *CachedPersistedRankingRepo {
	cachedPersistedRankingRepo := new(CachedPersistedRankingRepo)
	cachedPersistedRankingRepo.dynamodbRepo = dynamodbRepo
	cachedPersistedRankingRepo.dataCache = datacache
	return cachedPersistedRankingRepo
}

// GetNflPersistedRanking function
func (cachedPersistedRankingRepo *CachedPersistedRankingRepo) GetNflPersistedRanking(eventIds []int) map[int]domain.PersistedRanking {
	var persistedRankingInterface interface{}
	var persistedRankingPtr *domain.PersistedRanking
	var found bool
	persistedRankingMap := make(map[int]domain.PersistedRanking)
	notFoundEventIds := make([]int, 0)

	for _, eventID := range eventIds {
		persistedRankingInterface, found = cachedPersistedRankingRepo.dataCache.Get(strconv.Itoa(eventID))

		if found {
			persistedRankingPtr = persistedRankingInterface.(*domain.PersistedRanking)
			persistedRankingMap[eventID] = *persistedRankingPtr
		} else {
			notFoundEventIds = append(notFoundEventIds, eventID)
		}
	}

	if len(notFoundEventIds) > 0 {
		persistedRankingRepo := cachedPersistedRankingRepo.dynamodbRepo.GetNflPersistedRanking(notFoundEventIds)

		for k, v := range persistedRankingRepo {
			persistedRankingMap[k] = v
			cachedPersistedRankingRepo.dataCache.Set(strconv.Itoa(k), &v, cache.DefaultExpiration)
		}
	}

	return persistedRankingMap
}
