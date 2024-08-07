package dom

import (
	"syscall/js"
)

// HTMLNode es un wrapper alrededor de js.Value
type HTMLNode struct {
	value js.Value
}

// Element crea una nueva instancia de HTMLNode para un elemento HTML
func Element(tagName string) HTMLNode {
	element := Document().value.Call("createElement", tagName)
	return HTMLNode{value: element}
}

func (node HTMLNode) Value() js.Value {
	return node.value
}

type CanvasRenderingContext2D struct {
	value js.Value
}

// ClearRect limpia una sección del canvas.
func (ctx *CanvasRenderingContext2D) ClearRect(x, y, width, height float64) {
	ctx.value.Call("clearRect", x, y, width, height)
}

// SetFillStyle establece el estilo de relleno.
func (ctx *CanvasRenderingContext2D) SetFillStyle(style string) {
	ctx.value.Set("fillStyle", style)
}

// FillRect dibuja un rectángulo lleno.
func (ctx *CanvasRenderingContext2D) FillRect(x, y, width, height float64) {
	ctx.value.Call("fillRect", x, y, width, height)
}

func (ctx CanvasRenderingContext2D) SetStrokeStyle(style string) {
	ctx.value.Set("strokeStyle", style)
}

func (ctx CanvasRenderingContext2D) SetLineWidth(width int) {
	ctx.value.Set("lineWidth", width)
}

func (ctx CanvasRenderingContext2D) BeginPath() {
	ctx.value.Call("beginPath")
}

func (ctx CanvasRenderingContext2D) MoveTo(x, y float64) {
	ctx.value.Call("moveTo", x, y)
}

func (ctx CanvasRenderingContext2D) LineTo(x, y float64) {
	ctx.value.Call("lineTo", x, y)
}

func (ctx CanvasRenderingContext2D) Stroke() {
	ctx.value.Call("stroke")
}

// SetWidth establece el ancho del canvas.
func (node HTMLNode) SetWidth(width int) {
	node.value.Set("width", width)
}

// SetHeight establece la altura del canvas.
func (node HTMLNode) SetHeight(height int) {
	node.value.Set("height", height)
}

func NewCanvas(width, height int) (HTMLNode, CanvasRenderingContext2D) {
	canvas := Element("canvas")
	canvas.SetWidth(width)
	canvas.SetHeight(height)

	context := CanvasRenderingContext2D{value: canvas.value.Call("getContext", "2d")}
	return canvas, context
}
