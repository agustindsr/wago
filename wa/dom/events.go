package dom

import "syscall/js"

const (
	eventClick = "click"

	// Mouse events
	eventMouseEnter = "mouseenter"
	eventMouseLeave = "mouseleave"
	eventMouseOver  = "mouseover"
	eventMouseMove  = "mousemove"

	// Keyboard events
	eventKeyDown  = "keydown"
	eventKeyPress = "keypress"
	eventKeyUp    = "keyup"

	// Scroll events
	eventScroll = "scroll"
)

func (gv HTMLNode) AddEventListener(event string, fn Func) HTMLNode {
	wrappedFn := js.FuncOf(fn)
	gv.value.Call("addEventListener", event, wrappedFn)
	return gv
}

func (gv HTMLNode) OnClick(fn Func) HTMLNode {
	return gv.AddEventListener(eventClick, fn)
}

func (gv HTMLNode) OnHover(fn Func) HTMLNode {
	return gv.OnMouseEnter(fn).OnMouseLeave(fn)
}

func (gv HTMLNode) OnMouseEnter(fn Func) HTMLNode {
	return gv.AddEventListener(eventMouseEnter, fn)
}

func (gv HTMLNode) OnMouseLeave(fn Func) HTMLNode {
	return gv.AddEventListener(eventMouseLeave, fn)
}

func (gv HTMLNode) OnMouseOver(fn Func) HTMLNode {
	return gv.AddEventListener(eventMouseOver, fn)
}

func (gv HTMLNode) OnMouseMove(fn Func) HTMLNode {
	return gv.AddEventListener(eventMouseMove, fn)
}

func (gv HTMLNode) OnKeyDown(fn Func) HTMLNode {
	return gv.AddEventListener(eventKeyDown, fn)
}

func (gv HTMLNode) OnKeyPress(fn Func) HTMLNode {
	return gv.AddEventListener(eventKeyPress, fn)
}

func (gv HTMLNode) OnKeyUp(fn Func) HTMLNode {
	return gv.AddEventListener(eventKeyUp, fn)
}

func (gv HTMLNode) OnScroll(fn Func) HTMLNode {
	return gv.AddEventListener(eventScroll, fn)
}
