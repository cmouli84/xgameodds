package infrastructure

import (
	"fmt"
	"sync"

	"github.com/cmouli84/xgameodds/domain"

	"github.com/PuerkitoBio/goquery"
)

// TeamTrendsHandler struct
type TeamTrendsHandler struct {
}

// atsTrendsNcaabBaseURL const
const atsTrendsNcaabBaseURL = "https://www.teamrankings.com/ncb/trends/ats_trends/"

// overUnderNcaabBaseURL const
const overUnderNcaabBaseURL = "https://www.teamrankings.com/ncb/trends/ou_trends/"

// atsTrendsNflBaseURL const
const atsTrendsNflBaseURL = "https://www.teamrankings.com/nfl/trends/ats_trends/"

// overUnderNflBaseURL const
const overUnderNflBaseURL = "https://www.teamrankings.com/nfl/trends/ou_trends/"

// teamRankings Query string const
var teamTrendsQueryStrings = map[string]string{
	"all":                "",
	"home":               "?sc=is_home",
	"away":               "?sc=is_away",
	"fav":                "?sc=is_fav",
	"homefav":            "?sc=is_home_fav",
	"awayfav":            "?sc=is_away_fav",
	"underdog":           "?sc=is_dog",
	"homeunderdog":       "?sc=is_home_dog",
	"awayunderdog":       "?sc=is_away_dog",
	"conferencegames":    "?sc=is_conference",
	"nonconferencegames": "?sc=non_conference",
	"divisiongames":      "?sc=is_division",
	"nondivisiongames":   "?sc=non_division",
}

// NewTeamTrendsHandler function
func NewTeamTrendsHandler() *TeamTrendsHandler {
	teamTrendsHandler := new(TeamTrendsHandler)
	return teamTrendsHandler
}

// GetNflTeamTrends function
func (handler *TeamTrendsHandler) GetNflTeamTrends() domain.TeamTrends {
	return handler.getTeamTrends(atsTrendsNflBaseURL, overUnderNflBaseURL)
}

// GetNcaabTeamTrends function
func (handler *TeamTrendsHandler) GetNcaabTeamTrends() domain.TeamTrends {
	return handler.getTeamTrends(atsTrendsNcaabBaseURL, overUnderNcaabBaseURL)
}

// getTeamTrends function
func (handler *TeamTrendsHandler) getTeamTrends(atsTrendsBaseURL, overUnderBaseURL string) domain.TeamTrends {
	teamTrends := domain.TeamTrends(make(map[string]domain.Trend))

	//line below is my question
	wg := sync.WaitGroup{}
	// Ensure all routines finish before returning
	defer wg.Wait()

	for key, val := range teamTrendsQueryStrings {
		wg.Add(1)
		go getTrend(atsTrendsBaseURL+val, "ATS", key, &teamTrends, &wg)
	}

	for key, val := range teamTrendsQueryStrings {
		wg.Add(1)
		go getTrend(overUnderBaseURL+val, "OU", key, &teamTrends, &wg)
	}

	return teamTrends
}

// getTrend function
func getTrend(url, trendType, situation string, teamTrends *domain.TeamTrends, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println(url)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc.Find("table").First().Find("tbody tr").Each(func(i int, row *goquery.Selection) {
		var teamName, trendValue, underValue string
		row.Find("td").Each(func(j int, col *goquery.Selection) {
			if j == 0 {
				teamName = col.Find("a").Text()
			}
			if j == 2 {
				trendValue = col.Text()
			}
			if j == 3 {
				underValue = col.Text()
			}
		})

		var situationType *domain.Situation
		if _, ok := (*teamTrends)[teamName]; !ok {
			teamTrend := domain.Trend{}
			teamTrend.Ats = &domain.Situation{}
			teamTrend.OverUnder = &domain.Situation{}
			(*teamTrends)[teamName] = teamTrend
		}

		if trendType == "ATS" {
			situationType = (*teamTrends)[teamName].Ats
		} else {
			situationType = (*teamTrends)[teamName].OverUnder
			trendValue += "/" + underValue
		}

		switch situation {
		case "all":
			situationType.All = trendValue
		case "home":
			situationType.Home = trendValue
		case "away":
			situationType.Away = trendValue
		case "fav":
			situationType.Favorite = trendValue
		case "homefav":
			situationType.HomeFavorite = trendValue
		case "awayfav":
			situationType.AwayFavorite = trendValue
		case "underdog":
			situationType.Underdog = trendValue
		case "homeunderdog":
			situationType.HomeUnderdog = trendValue
		case "awayunderdog":
			situationType.AwayUnderdog = trendValue
		case "conferencegames":
			situationType.ConferenceGames = trendValue
		case "nonconferencegames":
			situationType.NonConferenceGames = trendValue
		case "divisiongames":
			situationType.DivisionGames = trendValue
		case "nondivisiongames":
			situationType.NonDivisionGames = trendValue
		}
	})

}
