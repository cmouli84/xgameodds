package interfaces

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/cmouli84/xgameodds/domain"
	"github.com/gorilla/mux"
)

// EventsInteractor interface
type EventsInteractor interface {
	GetEventsByDate(eventDate string) []domain.Event
}

// WebAPIHandler struct
type WebAPIHandler struct {
	EventsInteractor EventsInteractor
}

// GetEventsByDate func
func (handler *WebAPIHandler) GetEventsByDate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	eventDate := vars["eventdate"]

	events := handler.EventsInteractor.GetEventsByDate(eventDate)

	eventsPayload, jsonErr := json.Marshal(&events)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	res.Header().Add("Content-Type", "application/json")
	res.Write(eventsPayload)
}
