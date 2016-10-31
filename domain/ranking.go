package domain

// SonnyMooreRepository interface
type SonnyMooreRepository interface {
	GetSonnyMooreNflRanking() map[string]float64
}
