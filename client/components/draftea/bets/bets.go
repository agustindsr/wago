package bets

import (
	"wasm/client/components/draftea/betsgames"
	"wasm/client/components/draftea/betsplayers"
	"wasm/client/components/draftea/ticket"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

type OSB struct {
	BetsPlayer dom.HTMLNode
	BetsTeam   dom.HTMLNode
	TicketDom  dom.HTMLNode
}

func New() OSB {
	ticket := ticket.New()
	betsTeam := betsgames.New(ticket.AddBet, ticket.IsSelected).Render()
	betsPlayer := betsplayers.New(ticket.AddBet, ticket.IsSelected).Render()

	return OSB{
		BetsPlayer: betsPlayer,
		BetsTeam:   betsTeam,
		TicketDom:  ticket.Render(),
	}
}

func (b OSB) Render() dom.HTMLNode {
	return dom.Div().Tailwind(tlw.Flex, tlw.HScreen, tlw.BgGray700).Child(
		dom.Div().Tailwind(tlw.Flex, tlw.FlexCol, tlw.Flex1, tlw.OverflowYAuto).Child(
			b.BetsTeam.Tailwind(tlw.Mb4),
			b.BetsPlayer,
		),
		b.TicketDom.Tailwind(tlw.W1_4, tlw.HFull, tlw.BgGray800, tlw.P4, tlw.OverflowYAuto),
	)
}
