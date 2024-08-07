package dom

import "syscall/js"

type Event struct {
	Target  HTMLNode
	jsEvent js.Value
}

type Func func(Event)

func (e Event) Get(key string) js.Value {
	return e.jsEvent.Get(key)
}

// KeyCode returns the keyCode of the event
func (e Event) KeyCode() int {
	return e.Get("keyCode").Int()
}

// PreventDefault prevents the default action for the event
func (e Event) PreventDefault() {
	e.jsEvent.Call("preventDefault")
}

// StopPropagation stops the propagation of the event
func (e Event) StopPropagation() {
	e.jsEvent.Call("stopPropagation")
}

func (e Event) IsKeyCodeEnter() bool {
	return e.KeyCode() == 13
}
