package ticket

import (
	"fmt"
	"strconv"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

type AddToTicketFn func(betID, oddID, name, img string, multiplier float64)

type Ticket struct {
	ticketBets      []ticketBet
	EntryAmount     float64
	WinAmount       float64
	ticketContainer dom.HTMLNode
	winAmountLabel  dom.HTMLNode
	multiplierLabel dom.HTMLNode
	entryAmount     dom.HTMLNode
	betsList        dom.HTMLNode
}

type ticketBet struct {
	BetID      string
	OddID      string
	Name       string
	Img        string
	Multiplier float64
}

func New() *Ticket {
	t := &Ticket{
		EntryAmount: 100,
	}
	t.ticketContainer = dom.Div().SetID("ticket-container").Tailwind(tlw.BgGray800, tlw.RoundedLg, tlw.ShadowLg, tlw.P6, tlw.MaxWSm, tlw.Mx4, tlw.Mb4)

	t.ticketContainer.Child(
		dom.H2("Ticket").Tailwind(tlw.TextWhite, tlw.FontBold, tlw.Text2Xl, tlw.Mb4),
	)
	t.betsList = dom.Div()
	t.ticketContainer.Child(t.betsList)

	t.multiplierLabel = dom.Div().Tailwind(tlw.TextWhite, tlw.FontBold, tlw.TextLg)
	t.ticketContainer.Child(t.multiplierLabel)

	t.winAmountLabel = dom.Div().Tailwind(tlw.TextWhite, tlw.FontBold, tlw.TextLg)
	t.ticketContainer.Child(t.winAmountLabel)

	t.entryAmount = dom.Input().SetType("number").SetPlaceholder("Enter amount").SetValue(t.EntryAmount).Tailwind(tlw.WFull, tlw.P2, tlw.Mb4).OnChange(t.OnChangeEntryAmount)
	t.ticketContainer.Child(t.entryAmount)

	return t
}

func (t *Ticket) AddBet(betID, oddID, name, img string, multiplier float64) {
	t.ticketBets = append(t.ticketBets, ticketBet{
		BetID:      betID,
		OddID:      oddID,
		Name:       name,
		Img:        img,
		Multiplier: multiplier,
	})

	// Render and append the new bet element
	betElement := t.RenderBet(ticketBet{
		BetID:      betID,
		OddID:      oddID,
		Name:       name,
		Img:        img,
		Multiplier: multiplier,
	})
	t.betsList.Child(betElement)

	// Update the amount and win amount labels
	t.updateAmounts()
}

func (t *Ticket) CalculateMultiplier() float64 {
	multiplier := 1.0
	for _, bet := range t.ticketBets {
		multiplier *= bet.Multiplier
	}
	return multiplier
}

func (t *Ticket) CalculateWinAmount() float64 {
	return t.EntryAmount * t.CalculateMultiplier()
}

func (t *Ticket) updateAmounts() {
	winAmount := t.CalculateWinAmount()
	t.winAmountLabel.SetInnerHTML(fmt.Sprintf("Potential Win: $%.2f", winAmount))
	t.multiplierLabel.SetInnerHTML(fmt.Sprintf("Multiplier: %.2fx", t.CalculateMultiplier()))
}

func (t *Ticket) Render() dom.HTMLNode {
	return t.ticketContainer
}

func (t *Ticket) RenderBet(bet ticketBet) dom.HTMLNode {
	betContainer := dom.Div().Tailwind(tlw.Flex, tlw.ItemsCenter, tlw.BgGray700, tlw.RoundedMd, tlw.P4, tlw.Mb4)

	if bet.Img != "" {
		imgElement := dom.Img(bet.Img).Tailwind(tlw.W10, tlw.H10, tlw.Mr4)
		betContainer.Child(imgElement)
	}

	infoContainer := dom.Div().Tailwind(tlw.Flex1)

	nameElement := dom.Div().Tailwind(tlw.TextWhite, tlw.FontBold).SetInnerHTML(bet.Name)
	infoContainer.Child(nameElement)

	multiplierElement := dom.Div().Tailwind(tlw.TextWhite, tlw.TextSm).SetInnerHTML(fmt.Sprintf("Multiplier: %.2fx", bet.Multiplier))
	infoContainer.Child(multiplierElement)

	betContainer.Child(infoContainer)

	return betContainer
}

func (t *Ticket) OnChangeEntryAmount(e dom.Event) {
	valueStr := e.Target.GetValue().String()
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	t.EntryAmount = value
	t.updateAmounts()
}

func (t *Ticket) IsSelected(betID, oddID string) bool {
	for _, bet := range t.ticketBets {
		if bet.BetID == betID && bet.OddID == oddID {
			return true
		}
	}
	return false
}
