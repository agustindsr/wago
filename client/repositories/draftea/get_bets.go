package draftea

import (
	"fmt"
	"time"
	"wasm/client/wa/dom"
	"wasm/client/wa/http"
)

type (
	Response struct {
		Cursor string    `json:"cursor"`
		Cards  []CardDTO `json:"cards"`
	}

	CardDTO struct {
		IsInternal bool       `json:"isInternal"`
		Sport      SportDTO   `json:"sport"`
		League     LeagueDTO  `json:"league"`
		Game       GameDTO    `json:"game"`
		Groups     []GroupDTO `json:"groups"`
	}

	GroupDTO struct {
		CardGroup
		Value   *float64    `json:"value,omitempty"`
		Options []OptionDTO `json:"options"`
	}

	SportDTO struct {
		ID         string `json:"id"`
		Identifier string `json:"identifier"`
		Image      string `json:"image"`
	}

	GameDTO struct {
		ID   string    `json:"id"`
		Date time.Time `json:"date"`

		VersusData     *VersusData     `json:"VersusData,omitempty"`
		AutoRacingData *AutoRacingData `json:"AutoRacingData,omitempty"`
		TennisData     *TennisData     `json:"TennisData,omitempty"`
	}

	VersusData struct {
		HomeTeam TeamDTO `json:"homeTeam,omitempty"`
		AwayTeam TeamDTO `json:"awayTeam,omitempty"`
	}

	TennisData struct {
		FirstPlayer  PlayerDTO `json:"firstPlayer"`
		SecondPlayer PlayerDTO `json:"secondPlayer"`
		Modality     string    `json:"modality,omitempty"`
		Country      string    `json:"country,omitempty"`
		City         string    `json:"city,omitempty"`
	}

	AutoRacingData struct {
		CompetitionName string `json:"competitionName,omitempty"`
	}

	TeamDTO struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Abbreviation string `json:"abbreviation"`
		Image        string `json:"image"`
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

	LeagueDTO struct {
		ID string `json:"id"`
	}

	OptionDTO struct {
		OddID      string  `json:"oddId"`
		BetID      string  `json:"betId"`
		Prediction string  `json:"prediction"`
		Multiplier float64 `json:"multiplier"`
	}
)

type (
	CardGroupKey string

	CardGroup struct {
		Key  CardGroupKey `json:"key"`
		Name string       `json:"name"`
	}
)

func GetBets() Response {
	token := "eyJraWQiOiJJeHAzSm1ETTdtMFNvYXpMWlk1UEluN1o5VVwvN3I2QWJUV1ArRkNyMkx5QT0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIwZjAxN2MzZS00NTcwLTRjYTctYjNjZS00ZjlhZTNlMmJkYjMiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0yLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMl9kTHQ3bXlLb0wiLCJjbGllbnRfaWQiOiI2djJkNnRvMGJpaWM4Ym4xOGptdTc5OG5ydSIsIm9yaWdpbl9qdGkiOiJlZWE1ZjQyMS1hZjAwLTQ1ZDEtOTE4MC04ZDU0OTRkMGM1M2UiLCJldmVudF9pZCI6ImNmZjUzNWM5LWVhODYtNDAzYy05N2E3LTUzZDg2ZmFjYTczNyIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3MjI4OTMyMTQsImV4cCI6MTcyMjg5NjgxNCwiaWF0IjoxNzIyODkzMjE0LCJqdGkiOiJiOWUyYzFlOS0xODlhLTQ1MWEtYjIyNy0xYTI3NTdkNGNjNWEiLCJ1c2VybmFtZSI6IjBmMDE3YzNlLTQ1NzAtNGNhNy1iM2NlLTRmOWFlM2UyYmRiMyJ9.hwtuYu0sVJOqPbtlNOvj1KgLSFaM_eksYOFb63Gau4dhOD_zqZjMd1lJYqE3PEcSo2qTrpv4SPEi5V1c9e-0_zHeb8q8kTd6c4ZbsXHsy5bu9JU-2Z19nisTQ-xcQWoxXjFvfxO_T_cIYfbAXfFkIvNFMl8xL8Q5UF1GvW7qsFSMwtO3sw-TY_XlG3yJzAdmqcIM32dlgBZ6r2mnc0tqksF0V3xZa-VFzP9eA-CoBj1DvBKd8fJPibBx_edm_peBmwG1BWiKHmXihkbAM6Gff10_gxVKNDO_dqhAiBGRozwcxwQbYQNhdI2TxcjP1FsS74TGpUbRyecBoPIkEisAVQ"
	response, err := http.FetchData[Response]("https://api.dev.draftea.com/osb/bets/bff/game-and-team-game",
		nil,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", token),
		})
	if err != nil {
		dom.Alert("Error fetching data:")
	}

	return response
}
