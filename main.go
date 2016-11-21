package main

import (
	"net/http"

	"time"

	"github.com/cmouli84/xgameodds/infrastructure"
	"github.com/cmouli84/xgameodds/interfaces"
	"github.com/cmouli84/xgameodds/usecases"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

func main() {
	httpHandler := infrastructure.NewHTTPClient()
	scoreAPIHandler := infrastructure.NewScoreAPIHandler(httpHandler)
	sonnyMooreHandler := infrastructure.NewSonnyMooreHandler(httpHandler)

	dynamodbClient := infrastructure.NewDynamoDbClient()
	dynamodbHandler := infrastructure.NewDynamoDbHandler(dynamodbClient)

	scoreAPIRepository := interfaces.NewScoreAPIRepo(scoreAPIHandler)
	sonnyMooreRepository := interfaces.NewSonnyMooreRepo(sonnyMooreHandler)
	dynamodbRepository := interfaces.NewDynamoDbRepo(dynamodbHandler)

	scoreAPICache := cache.New(time.Minute*15, time.Minute*1)
	sonnyMooreCache := cache.New(time.Hour*6, time.Minute*1)
	dynamodbRepositoryCache := cache.New(time.Hour*24*7*30, time.Minute*1)

	cachedScoreAPIRepo := interfaces.NewCachedScoreAPIRepo(scoreAPIRepository, scoreAPICache)
	cachedSonnyMooreRepo := interfaces.NewCachedSonnyMooreRepo(sonnyMooreRepository, sonnyMooreCache)
	cachedDynamoDbRepo := interfaces.NewCachedPersistedRankingRepo(dynamodbRepository, dynamodbRepositoryCache)

	eventsInteractor := new(usecases.EventsInteractor)
	eventsInteractor.EventsRepository = cachedScoreAPIRepo
	eventsInteractor.SonnyMooreRepository = cachedSonnyMooreRepo
	eventsInteractor.DynamoDbRepository = cachedDynamoDbRepo

	webapiHandler := new(interfaces.WebAPIHandler)
	webapiHandler.EventsInteractor = eventsInteractor

	r := mux.NewRouter()
	r.HandleFunc("/api/nfl/events/{eventdate}", webapiHandler.GetNflEventsByDate).Methods("GET")

	http.ListenAndServe(":8181", r)
}
