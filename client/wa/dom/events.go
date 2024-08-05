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

func (node HTMLNode) AddEventListener(event string, fn Func) HTMLNode {
	wrappedFn := js.FuncOf(func(this js.Value, p []js.Value) any {
		e := Event{Target: node, jsEvent: p[0]}
		fn(e)
		return nil
	})
	node.value.Call("addEventListener", event, wrappedFn)
	return node
}

func (node HTMLNode) OnClick(fn Func) HTMLNode {
	return node.AddEventListener(eventClick, fn)
}

func (node HTMLNode) OnHover(fn Func) HTMLNode {
	return node.OnMouseEnter(fn).OnMouseLeave(fn)
}

func (node HTMLNode) OnMouseEnter(fn Func) HTMLNode {
	return node.AddEventListener(eventMouseEnter, fn)
}

func (node HTMLNode) OnMouseLeave(fn Func) HTMLNode {
	return node.AddEventListener(eventMouseLeave, fn)
}

func (node HTMLNode) OnMouseOver(fn Func) HTMLNode {
	return node.AddEventListener(eventMouseOver, fn)
}

func (node HTMLNode) OnMouseMove(fn Func) HTMLNode {
	return node.AddEventListener(eventMouseMove, fn)
}

func (node HTMLNode) OnKeyDown(fn Func) HTMLNode {
	return node.AddEventListener(eventKeyDown, fn)
}

func (node HTMLNode) OnKeyPress(fn Func) HTMLNode {
	return node.AddEventListener(eventKeyPress, fn)
}

func (node HTMLNode) OnKeyUp(fn Func) HTMLNode {
	return node.AddEventListener(eventKeyUp, fn)
}

func (node HTMLNode) OnScroll(fn Func) HTMLNode {
	return node.AddEventListener(eventScroll, fn)
}
