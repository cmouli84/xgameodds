package interfaces

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/cmouli84/xgameodds/domain"
	"github.com/gorilla/mux"
	"time"
	"log"
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

	startTime := time.Now()

	events := handler.EventsInteractor.GetNflEventsByDate(eventDate)

	eventsPayload, jsonErr := json.Marshal(&events)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	res.Header().Add("Content-Type", "application/json")
	res.Write(eventsPayload)
	log.Printf("Total Time taken for HTTP request %s/%s: %d", "nflevents", eventDate, time.Now().Sub(startTime)*time.Millisecond)
}

// GetNcaabEventsByDate func
func (handler *WebAPIHandler) GetNcaabEventsByDate(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	eventDate := vars["eventdate"]

	startTime := time.Now()

	events := handler.EventsInteractor.GetNcaabEventsByDate(eventDate)

	eventsPayload, jsonErr := json.Marshal(&events)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	res.Header().Add("Content-Type", "application/json")
	res.Write(eventsPayload)
	log.Printf("Total Time taken for HTTP request %s/%s: %d", "ncaabevents", eventDate, time.Now().Sub(startTime)*time.Millisecond)
}
