package counter

import (
	"fmt"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

type Counter struct {
	CountLabel      dom.HTMLNode
	IncrementButton dom.HTMLNode
	DecrementButton dom.HTMLNode
	ResetButton     dom.HTMLNode
	Count           int
}

func NewCounter() *Counter {
	c := &Counter{
		Count: 0,
	}

	c.CountLabel = dom.Span("0").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4)

	c.IncrementButton = dom.Button("Increment").OnClick(c.Increment).
		Tailwind(tlw.P2, tlw.BgBlue500, tlw.TextWhite, tlw.Rounded, tlw.HoverBgBlue900)
	c.DecrementButton = dom.Button("Decrement").OnClick(c.Decrement).
		Tailwind(tlw.P2, tlw.BgRed700, tlw.TextWhite, tlw.Rounded, tlw.HoverBgRed900)
	c.ResetButton = dom.Button("Reset").OnClick(c.Reset).
		Tailwind(tlw.P2, tlw.TextBlack, tlw.Border, tlw.Rounded, tlw.HoverTextWhite, tlw.HoverBgBlue700)

	return c
}

func (c *Counter) Render() dom.HTMLNode {
	container := dom.Div().
		Tailwind(tlw.MaxWXl, tlw.P4, tlw.RoundedLg, tlw.Flex, tlw.FlexCol).
		Child(
			dom.H2("Counter").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			c.CountLabel,
			dom.Div().Tailwind(tlw.Flex, tlw.SpaceX2).
				Child(
					c.IncrementButton,
					c.DecrementButton,
					c.ResetButton,
				),
		)

	return container
}

func (c *Counter) Increment(_ dom.Event) {
	c.Count++
	c.Update()
}

func (c *Counter) Decrement(_ dom.Event) {
	c.Count--
	c.Update()
}

func (c *Counter) Reset(_ dom.Event) {
	c.Count = 0
	c.Update()
}

func (c *Counter) Update() {
	c.CountLabel.SetInnerHTML(fmt.Sprintf("%d", c.Count))
}
