package usescases

import (
	"encoding/json"

	"github.com/cmouli84/xgameodds/domain"
	"github.com/cmouli84/xgameodds/infrastructure"
)

const scheduleUrl string = "http://api.thescore.com/nfl/schedule"

func GetEventsByDate(eventDate string) []domain.Event {
	scheduleResponse := infrastructure.GetHTTPResponse(scheduleUrl)

	schedule := domain.Schedule{}

	jsonerr := json.Unmarshal(scheduleResponse, &schedule)

	if jsonerr != nil {
		panic(jsonerr)
	}

    for i, group := range schedule.CurrentSeason {
        if (group.StartDate)
    }
}
