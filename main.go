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
	sonnyMooreHandler := infrastructure.NewSonnyMooreHandler(httpHandler)

	scoreAPIRepository := interfaces.NewScoreAPIRepo(scoreAPIHandler)
	sonnyMooreRepository := interfaces.NewSonnyMooreRepo(sonnyMooreHandler)

	eventsInteractor := new(usecases.EventsInteractor)
	eventsInteractor.EventsRepository = scoreAPIRepository
	eventsInteractor.SonnyMooreRepository = sonnyMooreRepository

	webapiHandler := new(interfaces.WebAPIHandler)
	webapiHandler.EventsInteractor = eventsInteractor

	r := mux.NewRouter()
	r.HandleFunc("/api/nfl/events/{eventdate}", webapiHandler.GetNflEventsByDate).Methods("GET")

	http.ListenAndServe(":8181", r)
}
