package interfaces

import (
	"github.com/cmouli84/xgameodds/domain"
	"github.com/patrickmn/go-cache"
)

const sonnyMooreCacheKey = "SONNY_RANKING"

// CachedSonnyMooreRepo struct
type CachedSonnyMooreRepo struct {
	sonnyMooreRepo domain.SonnyMooreRepository
	nflDataCache   *cache.Cache
	ncaabDataCache *cache.Cache
}

type getSonnyMooreRanking func() map[string]float64

// NewCachedSonnyMooreRepo function
func NewCachedSonnyMooreRepo(sonnyMooreInterface domain.SonnyMooreRepository, nflDataCache *cache.Cache, ncaabDataCache *cache.Cache) *CachedSonnyMooreRepo {
	cachedSonnyMooreRepo := new(CachedSonnyMooreRepo)
	cachedSonnyMooreRepo.sonnyMooreRepo = sonnyMooreInterface
	cachedSonnyMooreRepo.nflDataCache = nflDataCache
	cachedSonnyMooreRepo.ncaabDataCache = ncaabDataCache
	return cachedSonnyMooreRepo
}

// GetSonnyMooreNflRanking function
func (cachedSonnyMooreRepo *CachedSonnyMooreRepo) GetSonnyMooreNflRanking() map[string]float64 {
	return cachedSonnyMooreRepo.getSonnyMooreRanking(cachedSonnyMooreRepo.sonnyMooreRepo.GetSonnyMooreNflRanking, cachedSonnyMooreRepo.nflDataCache)
}

// GetSonnyMooreNcaabRanking function
func (cachedSonnyMooreRepo *CachedSonnyMooreRepo) GetSonnyMooreNcaabRanking() map[string]float64 {
	return cachedSonnyMooreRepo.getSonnyMooreRanking(cachedSonnyMooreRepo.sonnyMooreRepo.GetSonnyMooreNcaabRanking, cachedSonnyMooreRepo.ncaabDataCache)
}

// getSonnyMooreRanking function
func (cachedSonnyMooreRepo *CachedSonnyMooreRepo) getSonnyMooreRanking(getSonnyMooreRankingFn getSonnyMooreRanking, dataCache *cache.Cache) map[string]float64 {
	var sonnyMooreInterface interface{}
	var rankingMap *map[string]float64
	var found bool

	sonnyMooreInterface, found = dataCache.Get(sonnyMooreCacheKey)

	if found {
		rankingMap = sonnyMooreInterface.(*map[string]float64)
	} else {
		rankingMapRepo := getSonnyMooreRankingFn()
		rankingMap = &rankingMapRepo

		dataCache.Set(sonnyMooreCacheKey, rankingMap, cache.DefaultExpiration)
	}
	return *rankingMap
}
