package canvas

import (
	"fmt"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
)

type CanvasComponent struct {
	canvas     dom.HTMLNode
	context    dom.CanvasRenderingContext2D
	isDrawing  bool
	currentX   float64
	currentY   float64
	tool       string
	color      string
	brushSize  int
	clearColor string
}

// NewCanvasComponent crea una nueva instancia de CanvasComponent.
func NewCanvasComponent() *CanvasComponent {
	canvas, context := dom.NewCanvas(800, 600)
	cc := &CanvasComponent{
		canvas:     canvas,
		context:    context,
		isDrawing:  false,
		tool:       "brush",
		color:      "black",
		brushSize:  5,
		clearColor: "white",
	}
	cc.addEventListeners()
	return cc
}

// addEventListeners agrega los eventos del ratón al canvas.
func (cc *CanvasComponent) addEventListeners() {
	cc.canvas.OnMouseDown(cc.startDrawing)
	cc.canvas.OnMouseMove(cc.draw)
	cc.canvas.OnMouseUp(cc.stopDrawing)
	cc.canvas.OnMouseOut(cc.stopDrawing)
}

// startDrawing inicia el dibujo en el canvas.
func (cc *CanvasComponent) startDrawing(event dom.Event) {
	cc.isDrawing = true
	cc.currentX, cc.currentY = event.OffsetX(), event.OffsetY()
}

// draw realiza el dibujo en el canvas.
func (cc *CanvasComponent) draw(event dom.Event) {
	if !cc.isDrawing {
		return
	}
	newX, newY := event.OffsetX(), event.OffsetY()
	cc.context.SetStrokeStyle(cc.color)
	cc.context.SetLineWidth(cc.brushSize)
	cc.context.BeginPath()
	cc.context.MoveTo(cc.currentX, cc.currentY)
	cc.context.LineTo(newX, newY)
	cc.context.Stroke()
	cc.currentX, cc.currentY = newX, newY
}

// stopDrawing detiene el dibujo en el canvas.
func (cc *CanvasComponent) stopDrawing(event dom.Event) {
	cc.isDrawing = false
}

// clearCanvas borra el contenido del canvas.
func (cc *CanvasComponent) clearCanvas(_ dom.Event) {
	cc.context.SetFillStyle(cc.clearColor)
	cc.context.ClearRect(0, 0, 800, 600)
	cc.context.FillRect(0, 0, 800, 600)
}

// changeColor cambia el color del pincel.
func (cc *CanvasComponent) changeColor(color string) dom.Func {
	return func(_ dom.Event) {
		cc.color = color
	}
}

// changeBrushSize cambia el tamaño del pincel.
func (cc *CanvasComponent) changeBrushSize(size int) dom.Func {
	return func(_ dom.Event) {
		cc.brushSize = size
	}
}

// Render renderiza el componente de canvas.
func (cc *CanvasComponent) Render() dom.HTMLNode {
	colors := []string{"black", "red", "green", "blue", "yellow"}
	sizes := []int{2, 5, 10, 15, 20}

	colorButtons := dom.Div().Tailwind(tlw.Flex, tlw.SpaceX2)
	for _, color := range colors {
		button := dom.Button(color).
			Tailwind(tlw.P2, tlw.TextBlack, tlw.BgGray200, tlw.RoundedMd, tlw.HoverBgGray400, tlw.Border, tlw.BorderGray400, tlw.Shadow).
			OnClick(cc.changeColor(color))
		colorButtons.Child(button)
	}

	sizeButtons := dom.Div().Tailwind(tlw.Flex, tlw.SpaceX2)
	for _, size := range sizes {
		button := dom.Button(fmt.Sprintf("%d", size)).
			Tailwind(tlw.P2, tlw.TextBlack, tlw.BgGray200, tlw.RoundedMd, tlw.HoverBgGray400, tlw.Border, tlw.BorderGray400, tlw.Shadow).
			OnClick(cc.changeBrushSize(size))
		sizeButtons.Child(button)
	}

	clearButton := dom.Button("Clear").
		Tailwind(tlw.P2, tlw.BgRed500, tlw.TextWhite, tlw.RoundedMd, tlw.HoverBgRed700, tlw.Border, tlw.BorderGray400, tlw.Shadow).
		OnClick(cc.clearCanvas)

	container := dom.Div().
		Tailwind(tlw.MxAuto, tlw.P4, tlw.ShadowMd, tlw.RoundedLg, tlw.BgWhite).
		Child(
			dom.H2("Canvas Drawing App").
				Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			dom.Div().Child(colorButtons).Tailwind(tlw.Flex, tlw.SpaceX2, tlw.Mb2),
			dom.Div().Child(sizeButtons).Tailwind(tlw.Flex, tlw.SpaceX2, tlw.Mb2),
			clearButton.Tailwind(tlw.Mb4),
			cc.canvas.Tailwind(tlw.Border, tlw.BorderGray400, tlw.ShadowLg),
		)

	return container
}
