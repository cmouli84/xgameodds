package main

import (
	"net/http"

	"github.com/cmouli84/xgameodds/infrastructure"
	"github.com/cmouli84/xgameodds/interfaces"
	"github.com/cmouli84/xgameodds/usecases"
	"github.com/gorilla/mux"
)

func main() {
	httpHandler := infrastructure.NewHTTPClient()
	scoreAPIHandler := infrastructure.NewScoreAPIHandler(httpHandler)

	scoreAPIRepository := interfaces.NewScoreAPIRepo(scoreAPIHandler)

	eventsInteractor := new(usecases.EventsInteractor)
	eventsInteractor.EventsRepository = scoreAPIRepository

	webapiHandler := new(interfaces.WebAPIHandler)
	webapiHandler.EventsInteractor = eventsInteractor

	r := mux.NewRouter()
	r.HandleFunc("/api/events/{eventdate}", webapiHandler.GetEventsByDate).Methods("GET")

	http.ListenAndServe(":8181", r)
}
