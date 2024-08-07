package footer

import "wasm/pkg/dom"

func Render() dom.HTMLNode {
	footer := dom.Footer()
	footer.AddClass("footer")
	footer.SetInnerHTML("© 2024 My Website")

	return footer
}
