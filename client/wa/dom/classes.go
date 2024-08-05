package dom

func (node HTMLNode) AddClass(classes ...string) HTMLNode {
	for _, class := range classes {
		node.value.Get("classList").Call("add", class)
	}
	return node
}

func (node HTMLNode) RemoveClass(class string) HTMLNode {
	node.value.Get("classList").Call("remove", class)
	return node
}

func (node HTMLNode) ToggleClass(className string) HTMLNode {
	node.value.Get("classList").Call("toggle", className)
	return node
}

func (node HTMLNode) HasClass(className string) bool {
	return node.value.Get("classList").Call("contains", className).Bool()
}
