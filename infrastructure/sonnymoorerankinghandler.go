package infrastructure

import (
	"strconv"
	"strings"
)

// SonnyMooreHTTPClientHandler struct
type SonnyMooreHTTPClientHandler struct {
	httpClientInterface HTTPClientInterface
}

// sonnyMooreNflRankingURL constant
const sonnyMooreNflRankingURL = "http://sonnymoorepowerratings.com/nfl-foot.htm"

const sonnyMooreNflRankingStartsWith = "<B>\r\n"

const sonnyMooreNflRankingEndsWith = "\r\n</H3>"

// sonnyMooreNflRankingURL constant
const sonnyMooreNcaabRankingURL = "http://sonnymoorepowerratings.com/m-basket.htm"

const sonnyMooreNcaabRankingStartsWith = "<B>\r\n"

const sonnyMooreNcaabRankingEndsWith = "\r\n<BR>"

const rowDelimiter = "\r\n"

const colDelimiter = "  "

const trimChar = " "

// NewSonnyMooreHandler function
func NewSonnyMooreHandler(httpClientInterface HTTPClientInterface) *SonnyMooreHTTPClientHandler {
	sonnyMooreHTTPClientHandler := new(SonnyMooreHTTPClientHandler)
	sonnyMooreHTTPClientHandler.httpClientInterface = httpClientInterface
	return sonnyMooreHTTPClientHandler
}

// GetSonnyMooreNcaabRanking function
func (handler *SonnyMooreHTTPClientHandler) GetSonnyMooreNcaabRanking() map[string]float64 {
	return handler.extractSonnyMooreRanking(sonnyMooreNcaabRankingURL, sonnyMooreNcaabRankingStartsWith, sonnyMooreNcaabRankingEndsWith)
}

// GetSonnyMooreNflRanking function
func (handler *SonnyMooreHTTPClientHandler) GetSonnyMooreNflRanking() map[string]float64 {
	return handler.extractSonnyMooreRanking(sonnyMooreNflRankingURL, sonnyMooreNflRankingStartsWith, sonnyMooreNflRankingEndsWith)
}

func (handler *SonnyMooreHTTPClientHandler) extractSonnyMooreRanking(url, startsWith, endsWith string) map[string]float64 {
	response := handler.httpClientInterface.GetHTTPResponse(url)

	responseString := string(response)
	extractedString := responseString[strings.Index(responseString, startsWith)+len(startsWith):]
	extractedString = extractedString[:strings.Index(extractedString, endsWith)]

	teamRanks := handler.extractTeamRanks(extractedString)

	return teamRanks
}

func (handler *SonnyMooreHTTPClientHandler) extractTeamRanks(extractedTeams string) map[string]float64 {
	teams := handler.extractFields(extractedTeams, rowDelimiter, colDelimiter)

	teamMap := make(map[string]float64)
	for _, team := range teams {
		teamName := team[0][strings.Index(team[0], trimChar)+len(trimChar) : len(team[0])]

		rank, parseerr := strconv.ParseFloat(team[5], 64)
		if parseerr != nil {
			continue
		}
		teamMap[teamName] = rank
	}

	return teamMap
}

func (handler *SonnyMooreHTTPClientHandler) extractFields(extractedTeams, rowDelimiter, colDelimiter string) [][]string {
	teamStats := strings.Split(extractedTeams, rowDelimiter)

	teams := make([][]string, 0)
	for _, teamStat := range teamStats {
		team := make([]string, 0)
		for strings.Index(teamStat, colDelimiter) >= 0 {
			teamStat = strings.Trim(teamStat, trimChar)

			index := strings.Index(teamStat, colDelimiter)
			if index == -1 {
				index = len(teamStat)
			}
			team = append(team, teamStat[0:index])

			teamStat = teamStat[index:len(teamStat)]
		}

		teams = append(teams, team)
	}

	return teams
}
