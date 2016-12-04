package infrastructure

import (
	"fmt"
	"strings"

	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cmouli84/xgameodds/domain"
)

const nflEventsTableName = "NflEvents"

const ncaabEventsTableName = "NcaabEvents"

const primaryKey = "EventId"

const eventIDField = "EventId"

const homeTeamRankingField = "HomeTeamRanking"

const awayTeamRankingField = "AwayTeamRanking"

// NewDynamoDbClient function
func NewDynamoDbClient() *dynamodb.DynamoDB {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return dynamodb.New(session)
}

// DynamoDbHandler struct
type DynamoDbHandler struct {
	dynamodbClient *dynamodb.DynamoDB
}

// NewDynamoDbHandler function
func NewDynamoDbHandler(dynamodbClient *dynamodb.DynamoDB) *DynamoDbHandler {
	dynamodbHandler := &DynamoDbHandler{dynamodbClient: dynamodbClient}
	return dynamodbHandler
}

// GetNflPersistedRanking func
func (dynamodbHandler *DynamoDbHandler) GetNflPersistedRanking(eventIds []int) map[int]domain.PersistedRanking {
	return dynamodbHandler.getPersistedRanking(eventIds, nflEventsTableName)
}

// GetNcaabPersistedRanking func
func (dynamodbHandler *DynamoDbHandler) GetNcaabPersistedRanking(eventIds []int) map[int]domain.PersistedRanking {
	return dynamodbHandler.getPersistedRanking(eventIds, ncaabEventsTableName)
}

// getPersistedRanking function
func (dynamodbHandler *DynamoDbHandler) getPersistedRanking(eventIds []int, tableName string) map[int]domain.PersistedRanking {
	keys := make([]map[string]*dynamodb.AttributeValue, 0)

	for i := 0; i < len(eventIds); i++ {
		keyMap := make(map[string]*dynamodb.AttributeValue)
		keyMap[primaryKey] = &dynamodb.AttributeValue{N: aws.String(strconv.Itoa(eventIds[i]))}

		keys = append(keys, keyMap)
	}

	tableMap := make(map[string]*dynamodb.KeysAndAttributes)
	tableMap[tableName] = &dynamodb.KeysAndAttributes{
		Keys:                 keys,
		ProjectionExpression: aws.String(strings.Join([]string{eventIDField, homeTeamRankingField, awayTeamRankingField}, ",")),
	}

	params := &dynamodb.BatchGetItemInput{
		RequestItems: tableMap,
	}

	response, err := dynamodbHandler.dynamodbClient.BatchGetItem(params)

	if err != nil {
		fmt.Println(err)
		return map[int]domain.PersistedRanking{}
	}

	persistedRankingMap := make(map[int]domain.PersistedRanking)

	for _, item := range response.Responses[tableName] {
		if (item[homeTeamRankingField] == nil) || (item[awayTeamRankingField] == nil) {
			continue
		}

		eventID, _ := strconv.Atoi(*item[eventIDField].N)
		persistedRanking := domain.PersistedRanking{}
		persistedRanking.HomeRanking, _ = strconv.ParseFloat(*item[homeTeamRankingField].N, 64)
		persistedRanking.AwayRanking, _ = strconv.ParseFloat(*item[awayTeamRankingField].N, 64)

		persistedRankingMap[eventID] = persistedRanking
	}

	return persistedRankingMap
}
