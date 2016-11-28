package interfaces

import (
	"strconv"

	"github.com/cmouli84/xgameodds/domain"
	"github.com/patrickmn/go-cache"
)

// CachedPersistedRankingRepo struct
type CachedPersistedRankingRepo struct {
	dynamodbRepo   domain.DynamoDbRepository
	nflDataCache   *cache.Cache
	ncaabDataCache *cache.Cache
}

type getPersistedRanking func(eventIds []int) map[int]domain.PersistedRanking

// NewCachedPersistedRankingRepo function
func NewCachedPersistedRankingRepo(dynamodbRepo domain.DynamoDbRepository, nflDataCache *cache.Cache, ncaabDataCache *cache.Cache) *CachedPersistedRankingRepo {
	cachedPersistedRankingRepo := new(CachedPersistedRankingRepo)
	cachedPersistedRankingRepo.dynamodbRepo = dynamodbRepo
	cachedPersistedRankingRepo.nflDataCache = nflDataCache
	cachedPersistedRankingRepo.ncaabDataCache = ncaabDataCache
	return cachedPersistedRankingRepo
}

// GetNcaabPersistedRanking function
func (cachedPersistedRankingRepo *CachedPersistedRankingRepo) GetNcaabPersistedRanking(eventIDs []int) map[int]domain.PersistedRanking {
	return cachedPersistedRankingRepo.getPersistedRanking(eventIDs, cachedPersistedRankingRepo.dynamodbRepo.GetNcaabPersistedRanking, cachedPersistedRankingRepo.ncaabDataCache)
}

// GetNflPersistedRanking function
func (cachedPersistedRankingRepo *CachedPersistedRankingRepo) GetNflPersistedRanking(eventIDs []int) map[int]domain.PersistedRanking {
	return cachedPersistedRankingRepo.getPersistedRanking(eventIDs, cachedPersistedRankingRepo.dynamodbRepo.GetNflPersistedRanking, cachedPersistedRankingRepo.nflDataCache)
}

func (cachedPersistedRankingRepo *CachedPersistedRankingRepo) getPersistedRanking(eventIds []int, getPersistedRankingFn getPersistedRanking, dataCache *cache.Cache) map[int]domain.PersistedRanking {
	var persistedRankingInterface interface{}
	var persistedRankingPtr domain.PersistedRanking
	var found bool
	persistedRankingMap := make(map[int]domain.PersistedRanking)
	notFoundEventIds := make([]int, 0)

	for _, eventID := range eventIds {
		persistedRankingInterface, found = dataCache.Get(strconv.Itoa(eventID))

		if found {
			persistedRankingPtr = persistedRankingInterface.(domain.PersistedRanking)
			persistedRankingMap[eventID] = persistedRankingPtr
		} else {
			notFoundEventIds = append(notFoundEventIds, eventID)
		}
	}
	//fmt.Println("From Cache", persistedRankingMap)

	if len(notFoundEventIds) > 0 {
		persistedRankingRepo := getPersistedRankingFn(notFoundEventIds)
		//fmt.Println("From Repo", persistedRankingRepo)

		for k, v := range persistedRankingRepo {
			persistedRankingMap[k] = v
			dataCache.Set(strconv.Itoa(k), persistedRankingRepo[k], cache.DefaultExpiration)

			//persistedRankingInterface, found = dataCache.Get(strconv.Itoa(k))
			//fmt.Println(k, persistedRankingInterface.(domain.PersistedRanking).HomeRanking, persistedRankingInterface.(domain.PersistedRanking).AwayRanking)
		}
	}

	//fmt.Println("Consolidated", persistedRankingMap)

	return persistedRankingMap
}
