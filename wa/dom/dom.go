package dom

import "syscall/js"

// HTMLNode es un wrapper alrededor de js.Value
type HTMLNode struct {
	value js.Value
}

// Element crea una nueva instancia de HTMLNode para un elemento HTML
func Element(tagName string) HTMLNode {
	element := Document().value.Call("createElement", tagName)
	return HTMLNode{value: element}
}
