package interfaces

// SonnyMooreInterface interface
type SonnyMooreInterface interface {
	GetSonnyMooreNflRanking() map[string]float64
	GetSonnyMooreNcaabRanking() map[string]float64
}

type TeamnameMappingInterface interface {
	GetNcaabTeamNames() map[string]string
}

// SonnyMooreRepo struct
type SonnyMooreRepo struct {
	sonnyMooreInterface      SonnyMooreInterface
	teamnameMappingInterface TeamnameMappingInterface
}

// NewSonnyMooreRepo function
func NewSonnyMooreRepo(sonnyMooreInterface SonnyMooreInterface, teamnameMappingInterface TeamnameMappingInterface) *SonnyMooreRepo {
	sonnyMooreRepo := new(SonnyMooreRepo)
	sonnyMooreRepo.sonnyMooreInterface = sonnyMooreInterface
	sonnyMooreRepo.teamnameMappingInterface = teamnameMappingInterface
	return sonnyMooreRepo
}

// GetSonnyMooreNflRanking function
func (sonnyMooreRepo *SonnyMooreRepo) GetSonnyMooreNflRanking() map[string]float64 {
	return sonnyMooreRepo.sonnyMooreInterface.GetSonnyMooreNflRanking()
}

// GetSonnyMooreNcaabRanking function
func (sonnyMooreRepo *SonnyMooreRepo) GetSonnyMooreNcaabRanking() map[string]float64 {
	rankingMap := sonnyMooreRepo.sonnyMooreInterface.GetSonnyMooreNcaabRanking()
	teamMap := sonnyMooreRepo.teamnameMappingInterface.GetNcaabTeamNames()
	scoreAPIRankingMap := make(map[string]float64)

	for k, v := range rankingMap {
		scoreAPIRankingMap[teamMap[k]] = v
	}

	return scoreAPIRankingMap
}
