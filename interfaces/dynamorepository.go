package interfaces

import "github.com/cmouli84/xgameodds/domain"

// DynamoDbInterface interface
type DynamoDbInterface interface {
	GetNflPersistedRanking(eventIds []int) map[int]domain.PersistedRanking
}

// DynamoDbRepo struct
type DynamoDbRepo struct {
	dynamodbInterface DynamoDbInterface
}

// NewDynamoDbRepo function
func NewDynamoDbRepo(dynamodbInterface DynamoDbInterface) *DynamoDbRepo {
	dynamodbRepo := new(DynamoDbRepo)
	dynamodbRepo.dynamodbInterface = dynamodbInterface
	return dynamodbRepo
}

// GetNflPersistedRanking function
func (dynamodbRepo *DynamoDbRepo) GetNflPersistedRanking(eventIds []int) map[int]domain.PersistedRanking {
	return dynamodbRepo.dynamodbInterface.GetNflPersistedRanking(eventIds)
}
