package infrastructure

// TeamTrends map
type TeamTrends map[string]Trend

// Trend struct
type Trend struct {
	Ats       *Situation
	OverUnder *Situation
}

// Situation struct
type Situation struct {
	All                string
	Home               string
	Away               string
	Favorite           string
	HomeFavorite       string
	AwayFavorite       string
	Underdog           string
	HomeUnderdog       string
	AwayUnderdog       string
	ConferenceGames    string
	NonConferenceGames string
	DivisionGames      string
	NonDivisionGames   string
}
