package dom

import "syscall/js"

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

func PushHistoryState(state any, title, url string) {
	Global("history").Call("pushState", state, title, url)
}

func AwaitPromise(promise js.Value) (<-chan js.Value, <-chan error) {
	success := make(chan js.Value)
	failure := make(chan error)

	then := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		success <- args[0]
		return nil
	})

	catch := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		failure <- js.Error{Value: args[0]}
		return nil
	})

	promise.Call("then", then).Call("catch", catch)

	return success, failure
}
