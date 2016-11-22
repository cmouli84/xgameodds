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
	teamDynamoClient := infrastructure.NewDynamoDbClient()
	teamnamedbHandler := infrastructure.NewTeamnameDbHandler(teamDynamoClient)

	scoreAPIRepository := interfaces.NewScoreAPIRepo(scoreAPIHandler)
	sonnyMooreRepository := interfaces.NewSonnyMooreRepo(sonnyMooreHandler, teamnamedbHandler)
	dynamodbRepository := interfaces.NewDynamoDbRepo(dynamodbHandler)

	nflScoreAPICache := cache.New(time.Minute*15, time.Minute*1)
	nflSonnyMooreCache := cache.New(time.Hour*6, time.Minute*1)
	nflDynamodbRepositoryCache := cache.New(time.Hour*24*7*30, time.Minute*1)

	ncaabScoreAPICache := cache.New(time.Minute*15, time.Minute*1)
	ncaabSonnyMooreCache := cache.New(time.Hour*6, time.Minute*1)
	ncaabDynamodbRepositoryCache := cache.New(time.Hour*24*7*30, time.Minute*1)

	cachedScoreAPIRepo := interfaces.NewCachedScoreAPIRepo(scoreAPIRepository, nflScoreAPICache, ncaabScoreAPICache)
	cachedSonnyMooreRepo := interfaces.NewCachedSonnyMooreRepo(sonnyMooreRepository, nflSonnyMooreCache, ncaabSonnyMooreCache)
	cachedDynamoDbRepo := interfaces.NewCachedPersistedRankingRepo(dynamodbRepository, nflDynamodbRepositoryCache, ncaabDynamodbRepositoryCache)

	eventsInteractor := new(usecases.EventsInteractor)
	eventsInteractor.EventsRepository = cachedScoreAPIRepo
	eventsInteractor.SonnyMooreRepository = cachedSonnyMooreRepo
	eventsInteractor.DynamoDbRepository = cachedDynamoDbRepo

	webapiHandler := new(interfaces.WebAPIHandler)
	webapiHandler.EventsInteractor = eventsInteractor

	r := mux.NewRouter()
	r.HandleFunc("/api/nfl/events/{eventdate}", webapiHandler.GetNflEventsByDate).Methods("GET")
	r.HandleFunc("/api/ncaab/events/{eventdate}", webapiHandler.GetNflEventsByDate).Methods("GET")

	http.ListenAndServe(":8181", r)
}
