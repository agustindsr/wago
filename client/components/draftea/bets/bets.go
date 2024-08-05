package bets

import (
	"fmt"
	"wasm/client/repositories/draftea"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

type Component struct {
	BetsData draftea.Response
}

func New() *Component {
	resp := draftea.GetBets()
	return &Component{
		BetsData: resp,
	}
}

func (c *Component) Render() dom.HTMLNode {
	container := dom.Div().Tailwind(tlw.Flex, tlw.FlexWrap, tlw.SpaceX4, tlw.SpaceY4)

	for _, card := range c.BetsData.Cards {
		cardNode := c.renderCard(card)
		container.Child(cardNode)
	}

	return container
}

func (c *Component) renderCard(card draftea.CardDTO) dom.HTMLNode {
	cardContainer := dom.Div().Tailwind(tlw.Border, tlw.BorderGray300, tlw.RoundedLg, tlw.ShadowMd, tlw.P4, tlw.W1_4, tlw.BgWhite)

	versus := ""
	if card.Game.VersusData != nil {
		versus = fmt.Sprintf("%s vs %s", card.Game.VersusData.HomeTeam.Name, card.Game.VersusData.AwayTeam.Name)
	} else if card.Game.TennisData != nil {
		versus = fmt.Sprintf("%s vs %s", card.Game.TennisData.FirstPlayer.FullName, card.Game.TennisData.SecondPlayer.FullName)
	} else if card.Game.AutoRacingData != nil {
		versus = card.Game.AutoRacingData.CompetitionName
	}

	cardContainer.Child(
		dom.H3(versus).Tailwind(tlw.TextLg, tlw.FontBold, tlw.Mb2),
		dom.P(card.Game.Date.String()).Tailwind(tlw.TextSm, tlw.TextGray600, tlw.Mb4),
	)

	for _, group := range card.Groups {
		groupNode := c.renderGroup(group)
		cardContainer.Child(groupNode)
	}

	return cardContainer
}

func (c *Component) renderGroup(group draftea.GroupDTO) dom.HTMLNode {
	groupContainer := dom.Div().Tailwind(tlw.Mb4)

	groupContainer.Child(
		dom.H4(group.Name).Tailwind(tlw.TextBase, tlw.FontBold, tlw.Mb2),
	)

	for _, option := range group.Options {
		optionNode := c.renderOption(option)
		groupContainer.Child(optionNode)
	}

	return groupContainer
}

func (c *Component) renderOption(option draftea.OptionDTO) dom.HTMLNode {
	return dom.Button(fmt.Sprintf("%s - %.2f", option.Prediction, option.Multiplier)).Tailwind(tlw.P2, tlw.BgBlue500, tlw.TextWhite, tlw.Rounded, tlw.Mb2, tlw.Mr2)
}
