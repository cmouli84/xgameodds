package domain

// Event struct
type Event struct {
	Colours struct {
		Away string `json:"away"`
		Home string `json:"home"`
	} `json:"colours"`
	EventStatus          string      `json:"event_status"`
	HasPlayByPlayRecords bool        `json:"has_play_by_play_records"`
	ID                   int         `json:"id"`
	GameDate             string      `json:"game_date"`
	GameType             string      `json:"game_type"`
	GameDescription      interface{} `json:"game_description"`
	Tba                  bool        `json:"tba"`
	UpdatedAt            string      `json:"updated_at"`
	APIURI               string      `json:"api_uri"`
	ResourceURI          string      `json:"resource_uri"`
	Status               string      `json:"status"`
	Preview              string      `json:"preview"`
	Recap                string      `json:"recap,omitempty"`
	EventDetails         []struct {
		Label      string `json:"label"`
		Content    string `json:"content"`
		Identifier string `json:"identifier"`
	} `json:"event_details"`
	TvListingsByCountryCode struct {
		Us []struct {
			ShortName string `json:"short_name"`
			LongName  string `json:"long_name"`
		} `json:"us"`
		Ca []struct {
			ShortName string `json:"short_name"`
			LongName  string `json:"long_name"`
		} `json:"ca"`
	} `json:"tv_listings_by_country_code"`
	AwayTeam Team        `json:"away_team"`
	HomeTeam Team        `json:"home_team"`
	TopMatch interface{} `json:"top_match"`
	League   struct {
		Localizations struct {
		} `json:"localizations"`
		DailyRolloverOffset float64 `json:"daily_rollover_offset"`
		DailyRolloverTime   string  `json:"daily_rollover_time"`
		DefaultSection      string  `json:"default_section"`
		FullName            string  `json:"full_name"`
		MediumName          string  `json:"medium_name"`
		SeasonType          string  `json:"season_type"`
		ShortName           string  `json:"short_name"`
		Slug                string  `json:"slug"`
		SportName           string  `json:"sport_name"`
		UpdatedAt           string  `json:"updated_at"`
		APIURI              string  `json:"api_uri"`
		ResourceURI         string  `json:"resource_uri"`
	} `json:"league"`
	RedZone               bool `json:"red_zone"`
	HasTeamTwitterHandles bool `json:"has_team_twitter_handles"`
	Standings             struct {
		Away Standing `json:"away"`
		Home Standing `json:"home"`
	} `json:"standings"`
	Odd struct {
		AwayOdd   string `json:"away_odd"`
		HomeOdd   string `json:"home_odd"`
		ID        int    `json:"id"`
		Line      string `json:"line"`
		OverUnder string `json:"over_under"`
		APIURI    string `json:"api_uri"`
		Closing   string `json:"closing"`
	} `json:"odd"`
	SubscribableAlerts []struct {
		Key     string `json:"key"`
		Display string `json:"display"`
		Default bool   `json:"default"`
	} `json:"subscribable_alerts"`
	BoxScore struct {
		ID            int  `json:"id"`
		HasStatistics bool `json:"has_statistics"`
		Progress      struct {
			ClockLabel         string `json:"clock_label"`
			String             string `json:"string"`
			Status             string `json:"status"`
			EventStatus        string `json:"event_status"`
			Segment            int    `json:"segment"`
			SegmentString      string `json:"segment_string"`
			SegmentDescription string `json:"segment_description"`
			Clock              string `json:"clock"`
			Overtime           bool   `json:"overtime"`
		} `json:"progress"`
		UpdatedAt string `json:"updated_at"`
		APIURI    string `json:"api_uri"`
		Score     struct {
			Home struct {
				Score int `json:"score"`
			} `json:"home"`
			Away struct {
				Score int `json:"score"`
			} `json:"away"`
			WinningTeam string `json:"winning_team"`
			LosingTeam  string `json:"losing_team"`
			TieGame     bool   `json:"tie_game"`
		} `json:"score"`
		Minutes           interface{} `json:"minutes"`
		SegmentNumber     int         `json:"segment_number"`
		Down              interface{} `json:"down"`
		FormattedDistance interface{} `json:"formatted_distance"`
		YardsFromGoal     interface{} `json:"yards_from_goal"`
		HomeTimeoutsLeft  interface{} `json:"home_timeouts_left"`
		AwayTimeoutsLeft  interface{} `json:"away_timeouts_left"`
		UnderReview       bool        `json:"under_review"`
		TeamInPossession  interface{} `json:"team_in_possession"`
	} `json:"box_score"`
	Bowl          interface{} `json:"bowl"`
	Important     interface{} `json:"important"`
	Location      string      `json:"location"`
	SeasonWeek    string      `json:"season_week"`
	Stadium       string      `json:"stadium"`
	TotalQuarters int         `json:"total_quarters"`
	Week          int         `json:"week"`
	DisplayFpi    bool        `json:"display_fpi"`
}

// Team struct
type Team struct {
	Abbreviation string `json:"abbreviation"`
	Colour1      string `json:"colour_1"`
	Colour2      string `json:"colour_2"`
	Division     string `json:"division"`
	FullName     string `json:"full_name"`
	SearchName   string `json:"search_name"`
	ID           int    `json:"id"`
	Location     string `json:"location"`
	Logos        struct {
		Large   string      `json:"large"`
		Small   string      `json:"small"`
		W72Xh72 string      `json:"w72xh72"`
		Tiny    string      `json:"tiny"`
		Facing  interface{} `json:"facing"`
	} `json:"logos"`
	MediumName        string `json:"medium_name"`
	ShortName         string `json:"short_name"`
	Conference        string `json:"conference"`
	HasInjuries       bool   `json:"has_injuries"`
	HasRosters        bool   `json:"has_rosters"`
	UpdatedAt         string `json:"updated_at"`
	SubscriptionCount int    `json:"subscription_count"`
	APIURI            string `json:"api_uri"`
	ResourceURI       string `json:"resource_uri"`
}

// Standing struct
type Standing struct {
	LastFiveGamesRecord    string      `json:"last_five_games_record"`
	ShortRecord            string      `json:"short_record"`
	Streak                 string      `json:"streak"`
	ShortAtsRecord         interface{} `json:"short_ats_record"`
	ShortAwayRecord        string      `json:"short_away_record"`
	ShortConferenceRecord  string      `json:"short_conference_record"`
	ShortDivisionRecord    string      `json:"short_division_record"`
	ShortHomeRecord        string      `json:"short_home_record"`
	ConferenceRank         interface{} `json:"conference_rank"`
	ConferenceSeed         int         `json:"conference_seed"`
	DivisionRank           int         `json:"division_rank"`
	DivisionSeed           interface{} `json:"division_seed"`
	FormattedRank          string      `json:"formatted_rank"`
	Conference             string      `json:"conference"`
	ConferenceAbbreviation interface{} `json:"conference_abbreviation"`
	Division               string      `json:"division"`
	Place                  int         `json:"place"`
	APIURI                 string      `json:"api_uri"`
	DivisionRanking        int         `json:"division_ranking"`
}
