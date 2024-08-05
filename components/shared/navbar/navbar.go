package navbar

import (
	"fmt"
	"syscall/js"
	"wasm/wa/dom"
)

func Render() dom.NavElement {
	nav := dom.Nav()

	items := []string{"Home", "About", "Contact"}

	for _, item := range items {
		navItem := dom.Anchor().
			OnClick(ShowMessage()).
			OnHover(HandleMouseOver())
		navItem.SetInnerText(item)
		nav.Child(navItem)
	}

	dom.AppendToBody(nav.HTMLNode)
	return nav
}

func ShowMessage() dom.Func {
	return func(this js.Value, p []js.Value) any {
		dom.Alert(fmt.Sprintf("click %v", this))
		return nil
	}
}

func HandleMouseOver() dom.Func {
	return func(this js.Value, p []js.Value) any {
		dom.ToGlobalValue(this).ToggleClass("hover")
		return nil
	}
}
