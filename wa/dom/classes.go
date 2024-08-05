package dom

// AddClass agrega una clase al elemento
func (gv HTMLNode) AddClass(classes ...string) HTMLNode {
	for _, class := range classes {
		gv.value.Get("classList").Call("add", class)
	}
	return gv
}

// RemoveClass elimina una clase del elemento
func (gv HTMLNode) RemoveClass(class string) HTMLNode {
	gv.value.Get("classList").Call("remove", class)
	return gv
}

// ToggleClass alterna una clase del elemento
func (gv HTMLNode) ToggleClass(className string) HTMLNode {
	gv.value.Get("classList").Call("toggle", className)
	return gv
}

// HasClass verifica si el elemento tiene una clase
func (gv HTMLNode) HasClass(className string) bool {
	return gv.value.Get("classList").Call("contains", className).Bool()
}
