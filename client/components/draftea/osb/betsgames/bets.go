package betsgames

import (
	"encoding/json"
	"fmt"
	"sync"
	"wasm/client/components/draftea/osb/ticket"
	"wasm/client/css"
	"wasm/client/repositories/draftea"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
	"wasm/pkg/signal"
	signal2 "wasm/pkg/signal"
)

type OptionsKey struct {
	BetID string
	OddID string
}

type BetsTeamGame struct {
	BetsData      draftea.Response
	Ticket        *signal2.Signal[*ticket.Ticket]
	OptionsMapRef map[OptionsKey]dom.HTMLNode
	mu            sync.RWMutex
}

func New(ticket *signal2.Signal[*ticket.Ticket]) *BetsTeamGame {
	resp := GetBetsFromJSON()

	b := &BetsTeamGame{
		BetsData:      resp,
		Ticket:        ticket,
		OptionsMapRef: make(map[OptionsKey]dom.HTMLNode),
	}

	signal.NewEffect(b.updateSelectedBets, ticket.ToSignalInterface())

	return b
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
	cardContainer := dom.Div().Tailwind(tlw.RoundedLg, tlw.ShadowLg, tlw.P3, tlw.MaxWSm, tlw.Mx4).AddClass(css.BgTeriary700)

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
	e := dom.Button("").OnClick(b.clickOdd(card, option)).
		Tailwind(tlw.TextWhite, tlw.Rounded, tlw.TextSm, tlw.HoverBgGray600, tlw.W32).AddClass(css.BgTeriary500).
		Child(
			dom.Div().Tailwind(tlw.TextGray400).SetInnerHTML(option.Prediction),
			dom.Div().Tailwind(tlw.FontBold).SetInnerHTML(fmt.Sprintf("%.2fx", option.Multiplier)),
		)

	b.OptionsMapRef[OptionsKey{BetID: option.BetID, OddID: option.OddID}] = e

	return e
}

func (b *BetsTeamGame) clickOdd(card draftea.CardDTO, option draftea.OptionDTO) dom.Func {
	return func(e dom.Event) {
		gameName := fmt.Sprintf("%s VS %s", card.Game.VersusData.HomeTeam.Name, card.Game.VersusData.AwayTeam.Name)
		b.Ticket.Get().AddBet(option.BetID, option.OddID, gameName, "", option.Multiplier)
	}
}

func (b *BetsTeamGame) updateSelectedBets() {
	b.mu.Lock()
	defer b.mu.Unlock()

	selectedClasses := []tlw.TailwindClass{tlw.BgBlue500}
	for key, option := range b.OptionsMapRef {
		if b.Ticket.Get().IsSelected(key.BetID, key.OddID) {
			option.Tailwind(selectedClasses...)
		} else {
			option.TailwindRemove(selectedClasses...)
		}
	}
}
