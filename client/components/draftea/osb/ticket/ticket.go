package ticket

import (
	"fmt"
	"strconv"
	"sync"
	"wasm/client/wa/dom"
	"wasm/client/wa/dom/signal"
	tlw "wasm/client/wa/dom/tailwind"
)

type ticketBet struct {
	BetID        string
	OddID        string
	Name         string
	Img          string
	Multiplier   float64
	TicketBetRef dom.HTMLNode
}

type Ticket struct {
	TicketBets         *signal.Signal[[]ticketBet]
	EntryAmount        *signal.Signal[float64]
	WinAmount          float64
	signal             *signal.Signal[*Ticket]
	winAmountLabelRef  dom.HTMLNode
	multiplierLabelRef dom.HTMLNode
	entryAmountRef     dom.HTMLNode
	betsListRef        dom.HTMLNode
	mu                 sync.RWMutex
}

func New() *Ticket {
	t := &Ticket{
		TicketBets:  signal.NewSignal([]ticketBet{}),
		EntryAmount: signal.NewSignal(float64(100)),
	}

	t.signal = signal.NewSignal(t)

	return t
}

func (t *Ticket) Render() dom.HTMLNode {
	title := dom.H2("Ticket").Tailwind(tlw.TextWhite, tlw.FontBold, tlw.Text2Xl, tlw.Mb4)
	winningAmountLabel := dom.Div().Tailwind(tlw.TextWhite, tlw.FontBold, tlw.TextLg)
	multiplierLabel := dom.Div().Tailwind(tlw.TextWhite, tlw.FontBold, tlw.TextLg)
	entryAmountLabel := dom.Input().SetType("number").SetPlaceholder("Enter amount").SetValue(t.EntryAmount.Get()).
		Tailwind(tlw.WFull, tlw.P2, tlw.Mb4).OnChange(t.OnChangeEntryAmount)
	betListContainer := dom.Div()

	ticketContainer := dom.Div().SetID("ticket-container").Tailwind(tlw.BgGray800, tlw.RoundedLg, tlw.ShadowLg, tlw.P6, tlw.MaxWSm, tlw.Mx4, tlw.Mb4).
		Child(title, betListContainer, multiplierLabel, winningAmountLabel, entryAmountLabel)

	t.entryAmountRef = entryAmountLabel
	t.winAmountLabelRef = winningAmountLabel
	t.multiplierLabelRef = multiplierLabel
	t.betsListRef = betListContainer

	signal.NewEffect(t.updateAmounts, t.TicketBets.ToSignalInterface(), t.EntryAmount.ToSignalInterface())

	return ticketContainer
}

func (t *Ticket) AddBet(betID, oddID, name, img string, multiplier float64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	newTicketBet := ticketBet{
		BetID:      betID,
		OddID:      oddID,
		Name:       name,
		Img:        img,
		Multiplier: multiplier,
	}

	ticketBetElement := t.RenderBet(newTicketBet)
	newTicketBet.TicketBetRef = ticketBetElement

	ticketBets := t.TicketBets.Get()
	ticketBets = append(ticketBets, newTicketBet)
	t.TicketBets.Set(ticketBets)

	t.betsListRef.Child(ticketBetElement)

	t.signal.Set(t)
}

func (t *Ticket) RemoveBet(betID, oddID string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	ticketBets := t.TicketBets.Get()
	for i, bet := range ticketBets {
		if bet.BetID == betID && bet.OddID == oddID {
			bet.TicketBetRef.Remove()
			ticketBets = append(ticketBets[:i], ticketBets[i+1:]...)
			break
		}
	}
	t.TicketBets.Set(ticketBets)

	t.signal.Set(t)
}

func (t *Ticket) CalculateMultiplier() float64 {
	multiplier := 1.0
	for _, bet := range t.TicketBets.Get() {
		multiplier *= bet.Multiplier
	}
	return multiplier
}

func (t *Ticket) CalculateWinAmount() float64 {
	return t.EntryAmount.Get() * t.CalculateMultiplier()
}

func (t *Ticket) updateAmounts() {
	winAmount := t.CalculateWinAmount()
	t.winAmountLabelRef.SetInnerHTML(fmt.Sprintf("Potential Win: $%.2f", winAmount))
	t.multiplierLabelRef.SetInnerHTML(fmt.Sprintf("Multiplier: %.2fx", t.CalculateMultiplier()))
}

func (t *Ticket) RenderBet(bet ticketBet) dom.HTMLNode {
	betContainer := dom.Div().Tailwind(tlw.Flex, tlw.ItemsCenter, tlw.BgGray700, tlw.RoundedMd, tlw.P4, tlw.Mb4).SetID(bet.BetID + bet.OddID)

	if bet.Img != "" {
		imgElement := dom.Img(bet.Img).Tailwind(tlw.W10, tlw.H10, tlw.Mr4)
		betContainer.Child(imgElement)
	}

	infoContainer := dom.Div().Tailwind(tlw.Flex1)

	nameElement := dom.Div().Tailwind(tlw.TextWhite, tlw.FontBold).SetInnerHTML(bet.Name)
	infoContainer.Child(nameElement)

	multiplierElement := dom.Div().Tailwind(tlw.TextWhite, tlw.TextSm).SetInnerHTML(fmt.Sprintf("Multiplier: %.2fx", bet.Multiplier))
	infoContainer.Child(multiplierElement)

	deleteButton := dom.Button("Remove").OnClick(func(_ dom.Event) {
		t.RemoveBet(bet.BetID, bet.OddID)
	}).Tailwind(tlw.BgRed500, tlw.TextWhite, tlw.Rounded, tlw.P2, tlw.Ml4)

	betContainer.Child(infoContainer, deleteButton)

	return betContainer
}

func (t *Ticket) OnChangeEntryAmount(e dom.Event) {
	valueStr := e.Target.GetValue().String()
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return
	}
	t.EntryAmount.Set(value)
}

func (t *Ticket) IsSelected(betID, oddID string) bool {
	for _, bet := range t.TicketBets.Get() {
		if bet.BetID == betID && bet.OddID == oddID {
			return true
		}
	}
	return false
}

func (t *Ticket) GetSignal() *signal.Signal[*Ticket] {
	return t.signal
}
