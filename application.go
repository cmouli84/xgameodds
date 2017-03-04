package main

import (
	"fmt"
	"time"

	"github.com/cmouli84/xgameodds/infrastructure"
)

func main() {
	/*	httpHandler := infrastructure.NewHTTPClient()
		scoreAPIHandler := infrastructure.NewScoreAPIHandler(httpHandler)
		sonnyMooreHandler := infrastructure.NewSonnyMooreHandler(httpHandler)

		dynamodbClient := infrastructure.NewDynamoDbClient()
		dynamodbHandler := infrastructure.NewDynamoDbHandler(dynamodbClient)
		teamDynamoClient := infrastructure.NewDynamoDbClient()
		teamnamedbHandler := infrastructure.NewTeamnameDbHandler(teamDynamoClient)
		teamStatsDynamoClient := infrastructure.NewDynamoDbClient()
		teamStatsHandler := infrastructure.NewTeamStatsHandler(teamStatsDynamoClient)

		teamMapCache := cache.New(time.Hour*24*7*30, time.Minute*1)
		cachedTeamNameMap := interfaces.NewCachedTeamNameMapRepo(teamnamedbHandler, teamMapCache)

		scoreAPIRepository := interfaces.NewScoreAPIRepo(scoreAPIHandler)
		sonnyMooreRepository := interfaces.NewSonnyMooreRepo(sonnyMooreHandler, cachedTeamNameMap)
		dynamodbRepository := interfaces.NewDynamoDbRepo(dynamodbHandler)
		teamStatsRepository := interfaces.NewScoreAPITeamRepo(teamStatsHandler)

		nflScoreAPICache := cache.New(time.Minute*15, time.Minute*1)
		nflSonnyMooreCache := cache.New(time.Hour*6, time.Minute*1)
		nflDynamodbRepositoryCache := cache.New(time.Hour*24*7*30, time.Minute*1)
		nflTeamStatsCache := cache.New(time.Hour*12, time.Minute*1)

		ncaabScoreAPICache := cache.New(time.Minute*15, time.Minute*1)
		ncaabSonnyMooreCache := cache.New(time.Hour*6, time.Minute*1)
		ncaabDynamodbRepositoryCache := cache.New(time.Hour*24*7*30, time.Minute*1)
		ncaabTeamStatsCache := cache.New(time.Hour*12, time.Minute*1)

		cachedScoreAPIRepo := interfaces.NewCachedScoreAPIRepo(scoreAPIRepository, nflScoreAPICache, ncaabScoreAPICache)
		cachedScoreAPITeamRepo := interfaces.NewCachedScoreAPITeamRepo(teamStatsRepository, nflTeamStatsCache, ncaabTeamStatsCache)
		cachedSonnyMooreRepo := interfaces.NewCachedSonnyMooreRepo(sonnyMooreRepository, nflSonnyMooreCache, ncaabSonnyMooreCache)
		cachedDynamoDbRepo := interfaces.NewCachedPersistedRankingRepo(dynamodbRepository, nflDynamodbRepositoryCache, ncaabDynamodbRepositoryCache)

		eventsInteractor := new(usecases.EventsInteractor)
		eventsInteractor.EventsRepository = cachedScoreAPIRepo
		eventsInteractor.TeamRepository = cachedScoreAPITeamRepo
		eventsInteractor.SonnyMooreRepository = cachedSonnyMooreRepo
		eventsInteractor.DynamoDbRepository = cachedDynamoDbRepo

		webapiHandler := new(interfaces.WebAPIHandler)
		webapiHandler.EventsInteractor = eventsInteractor

		r := mux.NewRouter()
		r.HandleFunc("/api/nfl/events/{eventdate}", webapiHandler.GetNflEventsByDate).Methods("GET")
		r.HandleFunc("/api/ncaab/events/{eventdate}", webapiHandler.GetNcaabEventsByDate).Methods("GET")
		r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

		port := os.Getenv("PORT")
		if port == "" {
			port = "5000"
		}

		http.ListenAndServe(":"+port, r) */

	fmt.Println(time.Now())
	teamRankings := infrastructure.NewTeamRankingsHandler()
	var result = teamRankings.GetNcaabTeamTrends()

	for k, v := range result {
		fmt.Println(k)
		s := v.Ats
		fmt.Println(s.All, s.Home, s.Away, s.Favorite, s.HomeFavorite, s.AwayFavorite, s.Underdog,
			s.HomeUnderdog, s.AwayUnderdog, s.ConferenceGames, s.NonConferenceGames, s.DivisionGames, s.NonDivisionGames)
		s = v.OverUnder
		fmt.Println(s.All, s.Home, s.Away, s.Favorite, s.HomeFavorite, s.AwayFavorite, s.Underdog,
			s.HomeUnderdog, s.AwayUnderdog, s.ConferenceGames, s.NonConferenceGames, s.DivisionGames, s.NonDivisionGames)
	}

	fmt.Println(time.Now())
}
