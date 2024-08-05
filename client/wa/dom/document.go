package dom

func Document() HTMLNode {
	return Global("document")
}

func Body() HTMLNode {
	return HTMLNode{value: Document().Get("body").value}
}

func ElementByID(id string) HTMLNode {
	return HTMLNode{value: Document().value.Call("getElementById", id)}
}

func ElementByClassName(className string) HTMLNode {
	return HTMLNode{value: Document().value.Call("getElementsByClassName", className)}
}

func ElementByTagName(tagName string) HTMLNode {
	return HTMLNode{value: Document().value.Call("getElementsByTagName", tagName)}
}

func ElementByQuerySelector(selector string) HTMLNode {
	return HTMLNode{value: Document().value.Call("querySelector", selector)}
}

func ElementByQuerySelectorAll(selector string) HTMLNode {
	return HTMLNode{value: Document().value.Call("querySelectorAll", selector)}
}

func AppendToBody(gv HTMLNode) {
	Body().Child(gv)
}

func InjectCSS(css string) {
	style := Element("style")
	style.SetInnerHTML(css)
	Head().Child(style)
}
