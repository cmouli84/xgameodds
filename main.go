package main

import (
	"net/http"

	"github.com/cmouli84/xgameodds/infrastructure"
	"github.com/cmouli84/xgameodds/interfaces"
	"github.com/gorilla/mux"
)

func main() {
	httpHandler := infrastructure.NewHTTPClient()
	scoreApiHandler := infrastructure.NewScoreAPIHandler(httpHandler)

	scoreApiRepository := interfaces.NewScoreApiRepo(scoreApiHandler)

	eventsInteractor := new(usecases.EventsInteractor)
	eventsInteractor.EventsRepository = scoreApiRepository

	webapiHandler := new(interfaces.WebAPIHandler)
	webapiHandler.EventsInteractor = eventsInteractor

	r := mux.NewRouter()
	r.HandleFunc("/api/events/{eventDate}", webapiHandler.GetEventsByDate).Methods("GET")

	http.ListenAndServe(":8181", r)
}
