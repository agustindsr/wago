package dom

// SetAttribute establece un atributo del elemento
func (gv HTMLNode) SetAttribute(name, value string) HTMLNode {
	gv.value.Call("setAttribute", name, value)
	return gv
}

// GetAttribute obtiene el valor de un atributo del elemento
func (gv HTMLNode) GetAttribute(name string) string {
	return gv.value.Call("getAttribute", name).String()
}

// RemoveAttribute elimina un atributo del elemento
func (gv HTMLNode) RemoveAttribute(name string) HTMLNode {
	gv.value.Call("removeAttribute", name)
	return gv
}
