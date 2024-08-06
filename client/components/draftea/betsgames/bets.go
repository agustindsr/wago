package betsgames

import (
	"encoding/json"
	"fmt"
	"wasm/client/components/draftea/ticket"
	"wasm/client/repositories/draftea"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

var bets = `{
  "cursor": "eyJMaW1pdCI6MTAsIlNraXAiOjEwfQ",
  "cards": [
    {
      "isInternal": true,
      "sport": {
        "id": "62d6fb3ca1c099e6563b0ab3",
        "identifier": "football",
        "image": "https://cdn.draftea.com/sport-images/NFL.svg"
      },
      "league": {
        "id": "62d700f1a1c099e6563b0ad4"
      },
      "game": {
        "id": "66a7c89e80c445d1f39609a2",
        "date": "2024-09-06T00:20:00Z",
        "versusData": {
          "homeTeam": {
            "id": "62d6ff61a1c099e6563b0ab4",
            "name": "Chiefs",
            "abbreviation": "KC",
            "image": "https://cdn.dev.draftea.com/team-transparent-shirts/chiefs.svg"
          },
          "awayTeam": {
            "id": "62d6ff61a1c099e6563b0abf",
            "name": "Ravens",
            "abbreviation": "BAL",
            "image": "https://cdn.dev.draftea.com/team-transparent-shirts/ravens.svg"
          }
        }
      },
      "groups": [
        {
          "key": "spread",
          "name": "Spread",
          "options": [
            {
              "oddId": "odd_hzjP3mNW4uU1ztvoH",
              "betId": "bet_NWwD2DDn77wq0Oaux",
              "prediction": "Chiefs -3.00",
              "multiplier": 1.91
            },
            {
              "oddId": "odd_zqB2vjEf2vHD8luoN",
              "betId": "bet_71ZpP8rFwJ2dG9Nnz",
              "prediction": "Ravens +3.00",
              "multiplier": 1.91
            }
          ]
        },
        {
          "key": "totals",
          "name": "Puntos totales",
          "value": 46.5,
          "options": [
            {
              "oddId": "odd_BavUbQaFQg5ae15YH",
              "betId": "bet_SgVB5AiTyAzV1cb0p",
              "prediction": "↑ Más",
              "multiplier": 1.87
            },
            {
              "oddId": "odd_DTQYUhfdosiaWaYEF",
              "betId": "bet_SgVB5AiTyAzV1cb0p",
              "prediction": "Menos ↓",
              "multiplier": 1.95
            }
          ]
        },
        {
          "key": "money_line",
          "name": "Money Line",
          "options": [
            {
              "oddId": "odd_6Q8IarsIgPNDrWr0L",
              "betId": "bet_cCoSDPSC06dyrmP6j",
              "prediction": "Chiefs",
              "multiplier": 1.61
            },
            {
              "oddId": "odd_ZfA2HQpD9uLw8ecrY",
              "betId": "bet_p6dfrW84e4yzCpCAY",
              "prediction": "Ravens",
              "multiplier": 2.38
            }
          ]
        }
      ]
    },
    {
      "isInternal": false,
      "sport": {
        "id": "62d6fb3ca1c099e6563b0ab3",
        "identifier": "football",
        "image": "https://cdn.draftea.com/sport-images/NFL.svg"
      },
      "league": {
        "id": "62d700f1a1c099e6563b0ad4"
      },
      "game": {
        "id": "66a7c89e80c445d1f39609a2",
        "date": "2024-09-06T00:20:00Z",
        "versusData": {
          "homeTeam": {
            "id": "62d6ff61a1c099e6563b0ab4",
            "name": "Chiefs",
            "abbreviation": "KC",
            "image": "https://cdn.dev.draftea.com/team-transparent-shirts/chiefs.svg"
          },
          "awayTeam": {
            "id": "62d6ff61a1c099e6563b0abf",
            "name": "Ravens",
            "abbreviation": "BAL",
            "image": "https://cdn.dev.draftea.com/team-transparent-shirts/ravens.svg"
          }
        }
      },
      "groups": [
        {
          "key": "spread",
          "name": "Spread",
          "options": [
            {
              "oddId": "odd_gw2GjWRQOr8vPubE2",
              "betId": "bet_NWwD2DDn77wq0Oaux",
              "prediction": "Chiefs -3.00",
              "multiplier": 1.91
            },
            {
              "oddId": "odd_tkA2B1KJ0Urb8pdIt",
              "betId": "bet_71ZpP8rFwJ2dG9Nnz",
              "prediction": "Ravens +3.00",
              "multiplier": 1.91
            }
          ]
        },
        {
          "key": "totals",
          "name": "Puntos totales",
          "value": 46.5,
          "options": [
            {
              "oddId": "odd_STVVDnM0YdYwXyWKg",
              "betId": "bet_SgVB5AiTyAzV1cb0p",
              "prediction": "↑ Más",
              "multiplier": 1.87
            },
            {
              "oddId": "odd_mbvFgedqdSasxGOk5",
              "betId": "bet_SgVB5AiTyAzV1cb0p",
              "prediction": "Menos ↓",
              "multiplier": 1.95
            }
          ]
        },
        {
          "key": "money_line",
          "name": "Money Line",
          "options": [
            {
              "oddId": "odd_EWLL0OSubtNaLax2E",
              "betId": "bet_cCoSDPSC06dyrmP6j",
              "prediction": "Chiefs",
              "multiplier": 1.61
            },
            {
              "oddId": "odd_Ai6RtQo7UOJuOfLVU",
              "betId": "bet_p6dfrW84e4yzCpCAY",
              "prediction": "Ravens",
              "multiplier": 2.38
            }
          ]
        }
      ]
    }
  ]
}`

type BetsTeamGame struct {
	BetsData      draftea.Response
	AddToTicketFn ticket.AddToTicketFn
	IsSelected    func(betID, oddID string) bool
}

func New(addToTicketFn ticket.AddToTicketFn, isSelected func(betID, oddID string) bool) *BetsTeamGame {
	resp := GetBetsFromJSON()
	return &BetsTeamGame{
		BetsData:      resp,
		AddToTicketFn: addToTicketFn,
		IsSelected:    isSelected,
	}
}

func GetBetsFromJSON() draftea.Response {
	var response draftea.Response
	err := json.Unmarshal([]byte(bets), &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	return response
}

func (b *BetsTeamGame) Render() dom.HTMLNode {
	container := dom.Div().Tailwind(tlw.Flex, tlw.FlexWrap, tlw.JustifyCenter, tlw.ItemsCenter, tlw.SpaceY4, tlw.P4)

	for _, card := range b.BetsData.Cards {
		cardElement := b.renderCard(card)
		container.Child(cardElement)
	}

	return container
}

func (b *BetsTeamGame) renderCard(card draftea.CardDTO) dom.HTMLNode {
	cardContainer := dom.Div().Tailwind(tlw.BgGray800, tlw.RoundedLg, tlw.ShadowLg, tlw.P3, tlw.MaxWSm, tlw.Mx4)

	gameInfo := dom.Div().Tailwind(tlw.Flex, tlw.JustifyBetween, tlw.ItemsCenter, tlw.Mb4).Child(
		dom.Div().Child(
			dom.Img(card.Game.VersusData.HomeTeam.Image).Tailwind(tlw.W8, tlw.H8, tlw.Mr2),
			dom.Span(card.Game.VersusData.HomeTeam.Name).Tailwind(tlw.TextWhite, tlw.FontBold),
		),
		dom.Span("VS").Tailwind(tlw.TextWhite, tlw.TextXl, tlw.FontBold, tlw.Mx2),
		dom.Div().Child(
			dom.Img(card.Game.VersusData.AwayTeam.Image).Tailwind(tlw.W8, tlw.H8, tlw.Mr2),
			dom.Span(card.Game.VersusData.AwayTeam.Name).Tailwind(tlw.TextWhite, tlw.FontBold),
		),
	)

	date := card.Game.Date.Format("02/01/2006 15:04")
	gameDate := dom.Div().Tailwind(tlw.TextGray400, tlw.Mb1, tlw.TextCenter).SetInnerHTML(fmt.Sprintf("%s", date))

	groupsContainer := dom.Div().Tailwind(tlw.Mt4)
	for _, group := range card.Groups {
		groupElement := b.renderGroup(card, group)
		groupsContainer.Child(groupElement)
	}

	cardContainer.Child(gameInfo, gameDate, groupsContainer)

	return cardContainer
}

func (b *BetsTeamGame) renderGroup(card draftea.CardDTO, group draftea.GroupDTO) dom.HTMLNode {
	groupContainer := dom.Div().Tailwind(tlw.Flex, tlw.ItemsCenter, tlw.Mt2)

	groupTitle := dom.Div().Tailwind(tlw.TextWhite, tlw.Mr4, tlw.W32).SetInnerHTML(group.Name)

	optionsContainer := dom.Div().Tailwind(tlw.Flex, tlw.SpaceX1)
	for _, option := range group.Options {
		optionElement := b.renderOption(card, option)
		optionsContainer.Child(optionElement)
	}

	groupContainer.Child(groupTitle, optionsContainer)

	return groupContainer
}

func (b *BetsTeamGame) renderOption(card draftea.CardDTO, option draftea.OptionDTO) dom.HTMLNode {
	classes := []tlw.TailwindClass{
		tlw.BgGray700, tlw.TextWhite, tlw.Rounded, tlw.TextSm, tlw.HoverBgGray600, tlw.W32,
	}

	if b.IsSelected(option.BetID, option.OddID) {
		classes = append(classes, tlw.BgBlue500, tlw.TextWhite)
	}

	return dom.Button("").OnClick(b.clickOdd(card, option)).
		Tailwind(classes...).
		Child(
			dom.Div().Tailwind(tlw.TextGray400).SetInnerHTML(option.Prediction),
			dom.Div().Tailwind(tlw.FontBold).SetInnerHTML(fmt.Sprintf("%.2fx", option.Multiplier)),
		)
}

func (b *BetsTeamGame) clickOdd(card draftea.CardDTO, option draftea.OptionDTO) dom.Func {
	return func(e dom.Event) {
		gameName := fmt.Sprintf("%s VS %s", card.Game.VersusData.HomeTeam.Name, card.Game.VersusData.AwayTeam.Name)
		b.AddToTicketFn(option.BetID, option.OddID, gameName, "", option.Multiplier)
	}
}
