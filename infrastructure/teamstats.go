package infrastructure

import (
	"fmt"

	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"time"
	"log"
)

// EventStat struct
type EventStat struct {
	HomeTeamName string
	AwayTeamName string
	HomeScore    int
	AwayScore    int
	EventDate    string
	HomeOdds     float64
}

// TeamStatsHandler struct
type TeamStatsHandler struct {
	dynamodbClient *dynamodb.DynamoDB
}

// NewTeamStatsHandler function
func NewTeamStatsHandler(dynamodbClient *dynamodb.DynamoDB) *TeamStatsHandler {
	teamStatsHandler := &TeamStatsHandler{dynamodbClient: dynamodbClient}
	return teamStatsHandler
}

// GetNcaabTeamStats function
func (teamStatsHandler *TeamStatsHandler) GetNcaabTeamStats() []EventStat {
	return teamStatsHandler.getTeamStats(ncaabEventsTableName)
}

// GetNflTeamStats function
func (teamStatsHandler *TeamStatsHandler) GetNflTeamStats() []EventStat {
	return teamStatsHandler.getTeamStats(nflEventsTableName)
}

// getTeamStats function
func (teamStatsHandler *TeamStatsHandler) getTeamStats(tableName string) []EventStat {
	startTime := time.Now()

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
		AttributesToGet: []*string{
			aws.String("HomeTeamName"),
			aws.String("AwayTeamName"),
			aws.String("HomeScore"),
			aws.String("AwayScore"),
			aws.String("EventDate"),
			aws.String("HomeOdds"),
		},
	}

	resp, dynamoerr := teamStatsHandler.dynamodbClient.Scan(params)
	if dynamoerr != nil {
		fmt.Println(dynamoerr.Error())
		return []EventStat{}
	}

	eventStats := make([]EventStat, len(resp.Items))

	for index, item := range resp.Items {
		if (item["HomeScore"] == nil) || (item["AwayScore"] == nil) || (item["HomeTeamName"] == nil) {
			continue
		}
		eventStats[index].HomeTeamName = *item["HomeTeamName"].S
		eventStats[index].AwayTeamName = *item["AwayTeamName"].S
		eventStats[index].HomeScore, _ = strconv.Atoi(*item["HomeScore"].N)
		eventStats[index].AwayScore, _ = strconv.Atoi(*item["AwayScore"].N)
		eventStats[index].EventDate = *item["EventDate"].S
		eventStats[index].HomeOdds, _ = strconv.ParseFloat(*item["HomeOdds"].N, 64)
	}

	log.Printf("Time taken for Dynamo call GetTeamStats for %s: %d", tableName, time.Now().Sub(startTime) / time.Millisecond)

	return eventStats
}
