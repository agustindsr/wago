package draftea

import (
	"syscall/js"
	"wasm/client/wa/dom"
	"wasm/client/wa/http"
)

type Sport struct {
	ID         string `json:"id"`
	Identifier string `json:"identifier"`
	Image      string `json:"image"`
}

type Team struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Image        string `json:"image"`
}

type Game struct {
	ID         string `json:"id"`
	Date       string `json:"date"`
	VersusData struct {
		HomeTeam Team `json:"homeTeam"`
		AwayTeam Team `json:"awayTeam"`
	} `json:"versusData"`
}

type Option struct {
	OddID      string  `json:"oddId"`
	BetID      string  `json:"betId"`
	Prediction string  `json:"prediction"`
	Multiplier float64 `json:"multiplier"`
}

type Group struct {
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Options []Option `json:"options"`
}

type Card struct {
	IsInternal bool  `json:"isInternal"`
	Sport      Sport `json:"sport"`
	League     struct {
		ID string `json:"id"`
	} `json:"league"`
	Game   Game    `json:"game"`
	Groups []Group `json:"groups"`
}

type Response struct {
	Cursor string `json:"cursor"`
	Cards  []Card `json:"cards"`
}

func GetBets() dom.Func {
	return func(this js.Value, p []js.Value) any {
		response, err := http.FetchData[Response]("https://api.dev.draftea.com/osb/bets/bff/game-and-team-game",
			nil,
			map[string]string{
				"Authorization": "Bearer eyJraWQiOiJJeHAzSm1ETTdtMFNvYXpMWlk1UEluN1o5VVwvN3I2QWJUV1ArRkNyMkx5QT0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIwZjAxN2MzZS00NTcwLTRjYTctYjNjZS00ZjlhZTNlMmJkYjMiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0yLmFtYXpvbmF3cy5jb21cL3VzLWVhc3QtMl9kTHQ3bXlLb0wiLCJjbGllbnRfaWQiOiI2djJkNnRvMGJpaWM4Ym4xOGptdTc5OG5ydSIsIm9yaWdpbl9qdGkiOiI0ODAyZmVkOS02NzM2LTRhYzYtODk5NC0yODFiM2I0ZDA5OGQiLCJldmVudF9pZCI6ImZkODQ1M2Q2LTJjNjgtNDAyNy1hNjQ4LTZhZmQ1YTY4MWEzYSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3MjI4MDM2NzQsImV4cCI6MTcyMjgwNzI3NCwiaWF0IjoxNzIyODAzNjc0LCJqdGkiOiJlNjBmMTEwMS02Mjc0LTQyM2EtYTE2MS1mMGNlZWJkNTEzMGEiLCJ1c2VybmFtZSI6IjBmMDE3YzNlLTQ1NzAtNGNhNy1iM2NlLTRmOWFlM2UyYmRiMyJ9.baj0xsOkzh5qg4dYsY74V9BjWBGOsZEoPYJFrg1tzSYC7DdqhU7Vo_J8vNPk69e9ry4JxxK4KUkMBknOqBgUhTgWe1ClKVd90a5NF1Xcp4Soe07BVyoiVVeAvZsv9jeawFTfH9WkdCPBJLyCOWNH2R0VEdf4eYIspGipLsk1Hy_hqzPQh5Uc67puRGn8X_W6bXWOVL7ZpA77Z_Zd6gYDVpTXUcMmabl359XSGCdwIr4iawRuZCFCYHRxY5q1OovbQCJMkHcZwRCpU_2EDThAVoDpkWMz3Mm5WPi3abd6Dgq1feyw7TGwOCu7yEN3vv1pzobxKw-5ZnR_NCV9kyHGAA",
			})
		if err != nil {
			dom.Alert("Error fetching data:")
			return nil
		}

		renderTable(response)
		return nil
	}
}

func renderTable(data Response) {
	table := dom.ElementByID("bets")

	// Clear existing table rows
	table.ClearChildren()

	for _, card := range data.Cards {
		homeTeam := card.Game.VersusData.HomeTeam.Name
		awayTeam := card.Game.VersusData.AwayTeam.Name
		date := card.Game.Date

		row := dom.TR()
		row.Child(dom.TD().SetInnerHTML(homeTeam))
		row.Child(dom.TD().SetInnerHTML(awayTeam))
		row.Child(dom.TD().SetInnerHTML(date))

		table.Child(row.HTMLNode)
	}
}
