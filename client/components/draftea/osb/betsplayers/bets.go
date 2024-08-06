package betsplayers

import (
	"encoding/json"
	"fmt"
	"sync"
	"wasm/client/components/draftea/osb/ticket"
	draftea "wasm/client/repositories/draftea/players"
	"wasm/client/wa/dom"
	"wasm/client/wa/dom/signal"
	tlw "wasm/client/wa/dom/tailwind"
)

type OptionsKey struct {
	BetID string
	OddID string
}

type BetsPlayer struct {
	BetsData      draftea.Output
	Ticket        *signal.Signal[*ticket.Ticket]
	OptionsMapRef map[OptionsKey]dom.HTMLNode
	mu            sync.RWMutex
}

func New(ticket *signal.Signal[*ticket.Ticket]) *BetsPlayer {
	resp := GetBetsFromJSON()

	b := &BetsPlayer{
		BetsData:      resp,
		Ticket:        ticket,
		OptionsMapRef: make(map[OptionsKey]dom.HTMLNode),
	}

	signal.NewEffect(b.updateSelectedBets, ticket.ToSignalInterface())

	return b
}

func GetBetsFromJSON() draftea.Output {
	var response draftea.Output
	err := json.Unmarshal([]byte(bets), &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}
	return response
}

func (b *BetsPlayer) Render() dom.HTMLNode {
	container := dom.Div().Tailwind(tlw.Flex, tlw.FlexWrap, tlw.SpaceX4, tlw.P4)

	for _, bet := range b.BetsData.Bets {
		betElement := b.renderBetCard(bet)
		container.Child(betElement)
	}

	return container
}

func (b *BetsPlayer) renderBetCard(bet draftea.BetPlayerGameDTO) dom.HTMLNode {
	cardContainer := dom.Div().Tailwind(tlw.BgGray800, tlw.RoundedLg, tlw.ShadowLg, tlw.MaxW1_4, tlw.P0, tlw.M0)

	playerInfo := dom.Div().Tailwind(tlw.TextWhite, tlw.Px0, tlw.PT2, tlw.JustifyCenter)

	if bet.Player.Image != "" {
		playerInfo.Child(dom.Img(bet.Player.Image).Tailwind(tlw.W16, tlw.H16, tlw.Mb2))
	}

	if bet.Player.FullName != "" {
		playerInfo.Child(dom.Div().Tailwind(tlw.TextLg, tlw.FontBold, tlw.Px4).SetInnerHTML(bet.Player.FullName))
	}

	if bet.Position.Name != "" && bet.Game.VersusData != nil {
		playerInfo.Child(dom.Div().Tailwind(tlw.TextSm, tlw.TextGray400, tlw.Px4).SetInnerHTML(
			fmt.Sprintf("%s • %s vs %s",
				bet.Position.Name,
				getAbbreviation(bet.Game.VersusData.HomeTeam),
				getAbbreviation(bet.Game.VersusData.AwayTeam),
			),
		))
	}

	if !bet.Game.Date.IsZero() {
		playerInfo.Child(dom.Div().Tailwind(tlw.TextSm, tlw.TextGray400, tlw.Px4).SetInnerHTML(bet.Game.Date.Format("02/01/2006 15:04")))
	}

	for _, card := range bet.Cards {
		if card.IsInternal {
			continue
		}
		cardElement := b.renderCard(bet, card)
		playerInfo.Child(cardElement)
	}

	cardContainer.Child(playerInfo)

	return cardContainer
}

func getAbbreviation(team *draftea.TeamDTO) string {
	if team != nil {
		return team.Abbreviation
	}
	return ""
}

func (b *BetsPlayer) renderCard(bet draftea.BetPlayerGameDTO, card draftea.CardDTO) dom.HTMLNode {
	cardContainer := dom.Div().Tailwind(tlw.BgGray700, tlw.RoundedMd, tlw.P0, tlw.M0)

	cardValue := dom.Div().Tailwind(tlw.TextWhite, tlw.FontBold, tlw.TextXl, tlw.M0).SetInnerHTML(fmt.Sprintf("%.1f Tiros", card.Value))
	cardContainer.Child(cardValue)

	optionsContainer := dom.Div().Tailwind(tlw.Flex, tlw.SpaceX2, tlw.P0, tlw.M0)
	for _, odd := range card.Odds {
		oddElement := b.renderOdd(bet, odd)
		optionsContainer.Child(oddElement)
	}

	cardContainer.Child(optionsContainer)

	return cardContainer
}

func (b *BetsPlayer) renderOdd(bet draftea.BetPlayerGameDTO, odd draftea.OddDTO) dom.HTMLNode {
	e := dom.Button("").OnClick(b.clickOdd(bet, odd)).
		Tailwind(tlw.BgGray700, tlw.TextWhite, tlw.Rounded, tlw.TextSm, tlw.HoverBgGray600, tlw.W20, tlw.JustifyCenter).
		Child(
			dom.Div().Tailwind(tlw.TextGray300, tlw.TextSm, tlw.P0, tlw.M0).SetInnerHTML(odd.Prediction),
			dom.Div().Tailwind(tlw.TextWhite, tlw.FontBold, tlw.TextSm, tlw.P0, tlw.M0).SetInnerHTML(fmt.Sprintf("%.2fx", odd.Multiplier)),
		)

	b.OptionsMapRef[OptionsKey{BetID: bet.ID, OddID: odd.ID}] = e

	return e
}

func (b *BetsPlayer) clickOdd(bet draftea.BetPlayerGameDTO, odd draftea.OddDTO) dom.Func {
	return func(e dom.Event) {
		b.Ticket.Get().AddBet(bet.ID, odd.ID, bet.Player.FullName, bet.Player.Image, odd.Multiplier)
	}
}

func (b *BetsPlayer) updateSelectedBets() {
	b.mu.Lock()
	defer b.mu.Unlock()

	selectedClasses := []tlw.TailwindClass{tlw.BgBlue500}
	for key, option := range b.OptionsMapRef {
		if b.Ticket.Get().IsSelected(key.BetID, key.OddID) {
			option.Tailwind(selectedClasses...)
		} else {
			option.TailwindRemove(tlw.BgBlue500)
		}
	}
}
