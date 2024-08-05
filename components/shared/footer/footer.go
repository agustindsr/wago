package footer

import "wasm/wa/dom"

func Render() dom.FooterElement {
	footer := dom.Footer()
	footer.AddClass("footer")
	footer.SetInnerHTML("© 2024 My Website")

	dom.AppendToBody(footer.HTMLNode)
	return footer
}
