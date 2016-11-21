package interfaces

import (
	"github.com/cmouli84/xgameodds/domain"
	"github.com/patrickmn/go-cache"
)

const sonnyMooreCacheKey = "SONNY_RANKING"

// CachedSonnyMooreRepo struct
type CachedSonnyMooreRepo struct {
	sonnyMooreRepo domain.SonnyMooreRepository
	dataCache      *cache.Cache
}

// NewCachedSonnyMooreRepo function
func NewCachedSonnyMooreRepo(sonnyMooreInterface domain.SonnyMooreRepository, datacache *cache.Cache) *CachedSonnyMooreRepo {
	cachedSonnyMooreRepo := new(CachedSonnyMooreRepo)
	cachedSonnyMooreRepo.sonnyMooreRepo = sonnyMooreInterface
	cachedSonnyMooreRepo.dataCache = datacache
	return cachedSonnyMooreRepo
}

// GetSonnyMooreNflRanking function
func (cachedSonnyMooreRepo *CachedSonnyMooreRepo) GetSonnyMooreNflRanking() map[string]float64 {
	var sonnyMooreInterface interface{}
	var rankingMap *map[string]float64
	var found bool

	sonnyMooreInterface, found = cachedSonnyMooreRepo.dataCache.Get(sonnyMooreCacheKey)

	if found {
		rankingMap = sonnyMooreInterface.(*map[string]float64)
	} else {
		rankingMapRepo := cachedSonnyMooreRepo.sonnyMooreRepo.GetSonnyMooreNflRanking()
		rankingMap = &rankingMapRepo

		cachedSonnyMooreRepo.dataCache.Set(sonnyMooreCacheKey, rankingMap, cache.DefaultExpiration)
	}
	return *rankingMap
}
