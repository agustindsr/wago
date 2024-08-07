package counter

import (
	"fmt"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
	"wasm/pkg/signal"
)

type CounterSignal struct {
	Count          *signal.Signal[int]
	CountLabelRef  dom.HTMLNode
	ParityLabelRef dom.HTMLNode
}

func NewCounterSignal() *CounterSignal {
	c := &CounterSignal{
		Count: signal.NewSignal(0),
	}

	c.CountLabelRef = dom.Span("").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4)
	c.ParityLabelRef = dom.Span("").Tailwind(tlw.TextLg, tlw.FontBold, tlw.Mt2)

	signal.NewEffect(c.UpdateDOM, c.Count.ToSignalInterface())

	return c
}

func (c *CounterSignal) Render() dom.HTMLNode {
	incrementButton := dom.Button("Increment").OnClick(c.Increment).
		Tailwind(tlw.P2, tlw.BgBlue500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgBlue900)
	decrementButton := dom.Button("Decrement").OnClick(c.Decrement).
		Tailwind(tlw.P2, tlw.BgRed700, tlw.TextWhite, tlw.Rounded, tlw.HoverBgRed900)
	resetButton := dom.Button("Reset").OnClick(c.Reset).
		Tailwind(tlw.P2, tlw.BgWhite, tlw.TextBlack, tlw.Border, tlw.Rounded, tlw.HoverTextWhite, tlw.HoverBgBlue700)

	container := dom.Div().
		Tailwind(tlw.MaxWXl, tlw.P4, tlw.RoundedLg, tlw.Flex, tlw.FlexCol).
		Child(
			dom.H2("Counter").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			c.CountLabelRef,
			c.ParityLabelRef,
			dom.Div().Tailwind(tlw.Flex, tlw.SpaceX2).
				Child(
					incrementButton,
					decrementButton,
					resetButton,
				))

	return container
}

func (c *CounterSignal) Increment(_ dom.Event) {
	c.Count.Set(c.Count.Get() + 1)
}

func (c *CounterSignal) Decrement(_ dom.Event) {
	c.Count.Set(c.Count.Get() - 1)
}

func (c *CounterSignal) Reset(_ dom.Event) {
	c.Count.Set(0)
}

func (c *CounterSignal) Parity() string {
	if c.IsEven() {
		return "even"
	} else {
		return "odd"
	}
}

func (c *CounterSignal) IsEven() bool {
	return c.Count.Get()%2 == 0
}

func (c *CounterSignal) UpdateDOM() {
	c.CountLabelRef.SetInnerHTML(fmt.Sprintf("%d", c.Count.Get()))
	c.ParityLabelRef.SetInnerHTML(c.Parity())
}
