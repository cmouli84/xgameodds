package infrastructure

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"time"
	"log"
)

const ncaabScoreAPITeamTableName = "ScoreApiTeams"

const ncaabSonnyMooreTeamName = "SonnyMooreTeamName"

const ncaabScoreAPITeamName = "ScoreApiTeamName"

const ncaabTeamTrendTeamName = "TeamTrendTeamName"

// TeamnameDbHandler struct
type TeamnameDbHandler struct {
	dynamodbClient *dynamodb.DynamoDB
}

// NewTeamnameDbHandler function
func NewTeamnameDbHandler(dynamodbClient *dynamodb.DynamoDB) *TeamnameDbHandler {
	teamnameDynamodbHandler := &TeamnameDbHandler{dynamodbClient: dynamodbClient}
	return teamnameDynamodbHandler
}

// GetNcaabTeamNames function
func (teamnameDbHandler *TeamnameDbHandler) GetNcaabTeamNames() (map[string]string, map[string]string) {
	startTime := time.Now()

	params := &dynamodb.ScanInput{
		TableName: aws.String(ncaabScoreAPITeamTableName),
		AttributesToGet: []*string{
			aws.String(ncaabSonnyMooreTeamName),
			aws.String(ncaabScoreAPITeamName),
			aws.String(ncaabTeamTrendTeamName),
		},
	}

	resp, dynamoerr := teamnameDbHandler.dynamodbClient.Scan(params)

	sonnyTeamMap := make(map[string]string)
	teamTrendMap := make(map[string]string)
	if dynamoerr != nil {
		fmt.Println(dynamoerr.Error())
		return sonnyTeamMap, teamTrendMap
	}

	for _, item := range resp.Items {
		sonnyMooreTeamName := item[ncaabSonnyMooreTeamName].S
		scoreAPITeamName := item[ncaabScoreAPITeamName].S
		teamTrendTeamName := item[ncaabTeamTrendTeamName].S

		sonnyTeamMap[*sonnyMooreTeamName] = *scoreAPITeamName
		teamTrendMap[*teamTrendTeamName] = *scoreAPITeamName
	}

	log.Printf("Time taken for Dynamo call %s: %d", "GetNcaabTeamNames", time.Now().Sub(startTime) / time.Millisecond)

	return sonnyTeamMap, teamTrendMap
}
