package interfaces

import "github.com/cmouli84/xgameodds/domain"

// DynamoDbInterface interface
type DynamoDbInterface interface {
	GetNflPersistedRanking(eventIds []int) map[int]domain.PersistedRanking
	GetNcaabPersistedRanking(eventIds []int) map[int]domain.PersistedRanking
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

// GetNcaabPersistedRanking function
func (dynamodbRepo *DynamoDbRepo) GetNcaabPersistedRanking(eventIds []int) map[int]domain.PersistedRanking {
	rankingMap := make(map[int]domain.PersistedRanking)

	for i := 0; i <= (len(eventIds)-1)/100; i++ {
		maxArray := ((i + 1) * 100)
		if maxArray > len(eventIds) {
			maxArray = len(eventIds)
		}
		filteredRanking := dynamodbRepo.dynamodbInterface.GetNcaabPersistedRanking(eventIds[i*100 : maxArray])

		for k, v := range filteredRanking {
			rankingMap[k] = v
		}
	}

	return rankingMap
}
