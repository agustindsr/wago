package osb

import (
	"wasm/client/components/draftea/osb/betsgames"
	"wasm/client/components/draftea/osb/betsplayers"
	"wasm/client/components/draftea/osb/ticket"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
)

type OSB struct {
}

func NewOSB() *OSB {
	return &OSB{}
}

func (o *OSB) Render() dom.HTMLNode {
	t := ticket.New()
	ticketSignal := t.GetSignal()
	betsTeamGame := betsgames.New(ticketSignal)
	betsPlayerGame := betsplayers.New(ticketSignal)

	return dom.Div().Tailwind(tlw.Flex, tlw.HScreen).
		Child(
			dom.Div().Tailwind(tlw.Flex, tlw.FlexCol, tlw.Flex1, tlw.OverflowYAuto).
				Child(
					betsTeamGame.Render().Tailwind(tlw.Mb4),
					betsPlayerGame.Render().Tailwind(tlw.Mb4),
				),
			t.Render().Tailwind(tlw.W1_4, tlw.HFull, tlw.P4, tlw.OverflowYAuto),
		)
}
