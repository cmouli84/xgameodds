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

// scoreAPINflScheduleURL constant
const scoreAPINflScheduleURL = "http://api.thescore.com/nfl/schedule"

// scoreAPINcaabScheduleURL constant
const scoreAPINcaabScheduleURL = "http://api.thescore.com/ncaab/schedule"

// scoreAPINflEventsURLFormat constant
const scoreAPINflEventsURLFormat = "http://api.thescore.com/nfl/events?id.in=%s"

// scoreAPINcaabEventsURLFormat constant
const scoreAPINcaabEventsURLFormat = "http://api.thescore.com/ncaab/events?id.in=%s"

// NewScoreAPIHandler function
func NewScoreAPIHandler(httpClientInterface HTTPClientInterface) *ScoreAPIHTTPClientHandler {
	scoreAPIHTTPClientHandler := new(ScoreAPIHTTPClientHandler)
	scoreAPIHTTPClientHandler.httpClientInterface = httpClientInterface
	return scoreAPIHTTPClientHandler
}

// GetNflSchedule function
func (handler *ScoreAPIHTTPClientHandler) GetNflSchedule() ScoreAPISchedule {
	return handler.getSchedule(scoreAPINflScheduleURL)
}

// GetNcaabSchedule function
func (handler *ScoreAPIHTTPClientHandler) GetNcaabSchedule() ScoreAPISchedule {
	return handler.getSchedule(scoreAPINcaabScheduleURL)
}

// GetNflEvents function
func (handler *ScoreAPIHTTPClientHandler) GetNflEvents(eventIds []int) []ScoreAPIEvent {
	return handler.getEvents(scoreAPINflEventsURLFormat, eventIds)
}

// GetNcaabEvents function
func (handler *ScoreAPIHTTPClientHandler) GetNcaabEvents(eventIds []int) []ScoreAPIEvent {
	return handler.getEvents(scoreAPINcaabEventsURLFormat, eventIds)
}

// getSchedule function
func (handler *ScoreAPIHTTPClientHandler) getSchedule(scheduleURL string) ScoreAPISchedule {
	response := handler.httpClientInterface.GetHTTPResponse(scheduleURL)

	schedule := ScoreAPISchedule{}
	err := json.Unmarshal(response, &schedule)
	if err != nil {
		fmt.Println(err)
	}

	return schedule
}

// getEvents function
func (handler *ScoreAPIHTTPClientHandler) getEvents(eventsURL string, eventIds []int) []ScoreAPIEvent {
	eventIDText := make([]string, len(eventIds))

	for i, val := range eventIds {
		eventIDText[i] = strconv.Itoa(val)
	}

	response := handler.httpClientInterface.GetHTTPResponse(
		fmt.Sprintf(eventsURL, strings.Join(eventIDText, ",")))

	events := make([]ScoreAPIEvent, 0)
	err := json.Unmarshal(response, &events)
	if err != nil {
		fmt.Println(err)
	}

	return events
}
