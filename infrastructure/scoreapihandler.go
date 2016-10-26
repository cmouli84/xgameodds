package infrastructure

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ScoreAPIHTTPClientHandler struct
type ScoreAPIHTTPClientHandler struct {
	httpClientInterface HTTPClientInterface
}

// scoreAPIScheduleURL constant
const scoreAPIScheduleURL = "http://api.thescore.com/nfl/schedule"

// scoreAPIEventsURLFormat constant
const scoreAPIEventsURLFormat = "api.thescore.com/nfl/events?id.in=%s"

// NewScoreAPIHandler function
func NewScoreAPIHandler(httpClientInterface HTTPClientInterface) *ScoreAPIHTTPClientHandler {
	scoreAPIHTTPClientHandler := new(ScoreAPIHTTPClientHandler)
	scoreAPIHTTPClientHandler.httpClientInterface = httpClientInterface
	return scoreAPIHTTPClientHandler
}

// GetNflSchedule function
func (handler *ScoreAPIHTTPClientHandler) GetNflSchedule() ScoreAPISchedule {
	response := handler.httpClientInterface.GetHTTPResponse(scoreAPIScheduleURL)

	schedule := ScoreAPISchedule{}
	err := json.Unmarshal(response, &schedule)
	if err != nil {
		panic(err)
	}

	return schedule
}

// GetNflEvents function
func (handler *ScoreAPIHTTPClientHandler) GetNflEvents(eventIds []int) []ScoreAPIEvent {
	eventIDText := make([]string, len(eventIds))

	for i, val := range eventIds {
		eventIDText[i] = strconv.Itoa(val)
	}

	response := handler.httpClientInterface.GetHTTPResponse(fmt.Sprintf(scoreAPIEventsURLFormat, strings.Join(eventIDText, ",")))

	events := make([]ScoreAPIEvent, 0)
	err := json.Unmarshal(response, &events)
	if err != nil {
		panic(err)
	}

	return events
}
