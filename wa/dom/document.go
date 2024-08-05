package dom

// Document obtiene el objeto document
func Document() HTMLNode {
	return Global("document")
}

func Body() HTMLNode {
	return HTMLNode{value: Document().Get("body").value}
}

// ElementByID obtiene un elemento del DOM por su ID
func ElementByID(id string) HTMLNode {
	return HTMLNode{value: Document().value.Call("getElementById", id)}
}

// ElementByClassName obtiene un elemento del DOM por su clase
func ElementByClassName(className string) HTMLNode {
	return HTMLNode{value: Document().value.Call("getElementsByClassName", className)}
}

// ElementByTagName obtiene un elemento del DOM por su etiqueta
func ElementByTagName(tagName string) HTMLNode {
	return HTMLNode{value: Document().value.Call("getElementsByTagName", tagName)}
}

// ElementByQuerySelector obtiene un elemento del DOM por su selector
func ElementByQuerySelector(selector string) HTMLNode {
	return HTMLNode{value: Document().value.Call("querySelector", selector)}
}

// ElementByQuerySelectorAll obtiene todos los elementos del DOM por su selector
func ElementByQuerySelectorAll(selector string) HTMLNode {
	return HTMLNode{value: Document().value.Call("querySelectorAll", selector)}
}

// AppendToBody agrega un elemento al body del documento
func AppendToBody(gv HTMLNode) {
	ConsoleLog(Body().value)
	Body().Child(gv)
}

// InjectCSS inyecta un bloque de CSS en el documento
func InjectCSS(css string) {
	style := Element("style")
	style.SetInnerHTML(css)
	Head().Child(style)
}
