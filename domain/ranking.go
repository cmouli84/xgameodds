package domain

// SonnyMooreRepository interface
type SonnyMooreRepository interface {
	GetSonnyMooreNflRanking() map[string]float64
}

// DynamoDbRepository interface
type DynamoDbRepository interface {
	GetNflPersistedRanking(eventIds []int) map[int]PersistedRanking
}
