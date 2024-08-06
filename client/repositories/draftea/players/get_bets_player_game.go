package players

import "time"

type (
	Output struct {
		Cursor string             `json:"cursor"`
		Bets   []BetPlayerGameDTO `json:"bets"`
	}

	BetPlayerGameDTO struct {
		ID                  string      `json:"id"`
		Status              string      `json:"status"`
		IsLiveDataAvailable bool        `json:"isLiveDataAvailable"`
		Priority            int         `json:"priority"`
		Sport               SportDTO    `json:"sport"`
		Game                GameDTO     `json:"game"`
		Player              PlayerDTO   `json:"player"`
		Position            PositionDTO `json:"position"`
		Team                *TeamDTO    `json:"team"`
		RuleSet             RuleSetDTO  `json:"ruleSet"`
		League              LeagueDTO   `json:"league"`
		Cards               []CardDTO   `json:"cards"`
	}

	CardDTO struct {
		IsInternal bool     `json:"isInternal"`
		Value      float64  `json:"value"`
		Odds       []OddDTO `json:"odds"`
	}

	SportDTO struct {
		ID         string `json:"id"`
		Identifier string `json:"identifier"`
		Image      string `json:"image"`
	}

	GameDTO struct {
		ID   string    `json:"id"`
		Date time.Time `json:"date"`

		VersusData     *versusData     `json:"versusData,omitempty"`
		AutoRacingData *autoRacingData `json:"autoRacingData,omitempty"`
		TennisData     *tennisData     `json:"tennisData,omitempty"`
	}

	versusData struct {
		HomeTeam *TeamDTO `json:"homeTeam,omitempty"`
		AwayTeam *TeamDTO `json:"awayTeam,omitempty"`
	}

	tennisData struct {
		FirstPlayer  *PlayerDTO `json:"firstPlayer,omitempty"`
		SecondPlayer *PlayerDTO `json:"secondPlayer,omitempty"`
		Modality     string     `json:"modality,omitempty"`
		Country      string     `json:"country,omitempty"`
		City         string     `json:"city,omitempty"`
	}

	autoRacingData struct {
		CompetitionName string `json:"competitionName,omitempty"`
	}

	TeamDTO struct {
		ID                  string `json:"id"`
		Name                string `json:"name"`
		Abbreviation        string `json:"abbreviation"`
		TransparentShirtURL string `json:"transparentShirtUrl"`
	}

	PositionDTO struct {
		ID         string `json:"id"`
		Identifier string `json:"identifier"`
		Name       string `json:"name"`
		SportID    string `json:"sportId"`
	}

	PlayerDTO struct {
		ID               string `json:"id"`
		FullName         string `json:"fullName"`
		FirstName        string `json:"firstName,omitempty"`
		LastName         string `json:"lastName,omitempty"`
		AbbreviationName string `json:"abbreviationName,omitempty"`
		Image            string `json:"image,omitempty"`
		CountryCode      string `json:"countryCode,omitempty"`
	}

	RuleSetDTO struct {
		Identifier  string `json:"identifier"`
		Description string `json:"description"`
	}

	LeagueDTO struct {
		ID string `json:"id"`
	}

	OddDTO struct {
		ID         string  `json:"id"`
		Prediction string  `json:"prediction"`
		Multiplier float64 `json:"multiplier"`
	}
)
