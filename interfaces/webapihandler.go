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
	GetNflEventsByDate(eventDate string) []domain.Event
	GetNcaabEventsByDate(eventDate string) []domain.Event
}

// WebAPIHandler struct
type WebAPIHandler struct {
	EventsInteractor EventsInteractor
}

// GetNflEventsByDate func
func (handler *WebAPIHandler) GetNflEventsByDate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	eventDate := vars["eventdate"]

	events := handler.EventsInteractor.GetNflEventsByDate(eventDate)

	eventsPayload, jsonErr := json.Marshal(&events)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	res.Header().Add("Content-Type", "application/json")
	res.Write(eventsPayload)
}

// GetNcaabEventsByDate func
func (handler *WebAPIHandler) GetNcaabEventsByDate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	eventDate := vars["eventdate"]

	events := handler.EventsInteractor.GetNcaabEventsByDate(eventDate)

	eventsPayload, jsonErr := json.Marshal(&events)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	res.Header().Add("Content-Type", "application/json")
	res.Write(eventsPayload)
}
