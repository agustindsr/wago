package dom

func (node HTMLNode) SetAttribute(name, value string) HTMLNode {
	node.value.Call("setAttribute", name, value)
	return node
}

func (node HTMLNode) GetAttribute(name string) string {
	return node.value.Call("getAttribute", name).String()
}

func (node HTMLNode) RemoveAttribute(name string) HTMLNode {
	node.value.Call("removeAttribute", name)
	return node
}

func (node HTMLNode) SetID(id string) HTMLNode {
	return node.SetAttribute("id", id)
}

func (node HTMLNode) GetID() string {
	return node.GetAttribute("id")
}

func (node HTMLNode) RemoveID() HTMLNode {
	return node.RemoveAttribute("id")
}

func (node HTMLNode) SetPlaceholder(placeholder string) HTMLNode {
	return node.SetAttribute("placeholder", placeholder)
}

func (node HTMLNode) GetPlaceholder() string {
	return node.GetAttribute("placeholder")
}
