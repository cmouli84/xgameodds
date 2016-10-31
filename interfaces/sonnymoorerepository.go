package interfaces

// SonnyMooreInterface interface
type SonnyMooreInterface interface {
	GetSonnyMooreNflRanking() map[string]float64
}

// SonnyMooreRepo struct
type SonnyMooreRepo struct {
	sonnyMooreInterface SonnyMooreInterface
}

// NewSonnyMooreRepo function
func NewSonnyMooreRepo(sonnyMooreInterface SonnyMooreInterface) *SonnyMooreRepo {
	sonnyMooreRepo := new(SonnyMooreRepo)
	sonnyMooreRepo.sonnyMooreInterface = sonnyMooreInterface
	return sonnyMooreRepo
}

// GetSonnyMooreNflRanking function
func (sonnyMooreRepo *SonnyMooreRepo) GetSonnyMooreNflRanking() map[string]float64 {
	return sonnyMooreRepo.sonnyMooreInterface.GetSonnyMooreNflRanking()
}
