package dom

import (
	"syscall/js"
	"wasm/client/wa/dom/tailwind"
)

const (
	tableElement  = "table"
	trElement     = "tr"
	tdElement     = "td"
	divElement    = "div"
	h1Element     = "h1"
	h2Element     = "h2"
	h3Element     = "h3"
	h4Element     = "h4"
	h5Element     = "h5"
	h6Element     = "h6"
	navElement    = "nav"
	footerElement = "footer"
	bodyElement   = "body"
	headElement   = "head"
	spanElement   = "span"
	pElement      = "p"
	brElement     = "br"
	strongElement = "strong"
	ulElement     = "ul"
	liElement     = "li"
	inputElement  = "input"
	buttonElement = "button"
)

func Anchor(t any) HTMLNode {
	return Element("a").SetInnerHTML(t)
}

func Div() HTMLNode {
	return Element(divElement)
}

func Span(t any) HTMLNode {
	return Element(spanElement).SetInnerHTML(t)
}

func P(t any) HTMLNode {
	return Element(pElement).SetInnerHTML(t)
}

func BR() HTMLNode {
	return Element(brElement)
}

func Strong() HTMLNode {
	return Element(strongElement)
}

func Table() HTMLNode {
	return Element(tableElement)
}

func TR() HTMLNode {
	return Element(trElement)
}

func TD() HTMLNode {
	return Element(tdElement)
}

func Footer() HTMLNode {
	return Element(footerElement)
}

func Nav() HTMLNode {
	return Element(navElement)
}

func H1(t any) HTMLNode {
	return Element(h1Element).SetInnerHTML(t)
}

func H2(t any) HTMLNode {
	return Element(h2Element).SetInnerHTML(t)
}

func H3(t any) HTMLNode {
	return Element(h3Element).SetInnerHTML(t)
}

func Img(src string) HTMLNode {
	return Element("img").Set("src", src)
}

func H4(t any) HTMLNode {
	return Element(h4Element).SetInnerHTML(t)
}

func H5() HTMLNode {
	return Element(h5Element)
}

func H6() HTMLNode {
	return Element(h6Element)
}

func Pre(t any) HTMLNode {
	return Element("pre").SetInnerHTML(t)
}

func Head() HTMLNode {
	return Element("head")
}

func Input() HTMLNode {
	return Element(inputElement)
}

func Button(t string) HTMLNode {
	return Element(buttonElement).SetInnerText(t)
}

func UL() HTMLNode {
	return Element(ulElement)
}

func LI() HTMLNode {
	return Element(liElement)
}

func (node HTMLNode) SetInnerHTML(html any) HTMLNode {
	node.value.Set("innerHTML", html)
	return node
}

func (node HTMLNode) SetInnerText(text string) HTMLNode {
	node.value.Set("innerText", text)
	return node
}

func (node HTMLNode) Child(child ...HTMLNode) HTMLNode {
	for _, c := range child {
		node.value.Call("appendChild", c.Value())
	}
	return node
}

func (node HTMLNode) RemoveChild(child HTMLNode) HTMLNode {
	node.value.Call("removeChild", child.Value())
	return node
}

func (node HTMLNode) ClearChildren() HTMLNode {
	node.SetInnerHTML("")
	return node
}

func (node HTMLNode) Get(name string) HTMLNode {
	return HTMLNode{value: node.value.Get(name)}
}

func (node HTMLNode) Index(i int) HTMLNode {
	return HTMLNode{value: node.value.Index(i)}
}

func (node HTMLNode) GetInnerText() string {
	return node.Get("innerText").String()
}

func (node HTMLNode) GetInnerHTML() string {
	return node.Get("innerHTML").String()
}

func (node HTMLNode) GetValue() HTMLNode {
	return node.Get("value")
}

func (node HTMLNode) Float() float64 {
	return node.value.Float()
}

func (node HTMLNode) Set(name string, value any) HTMLNode {
	node.value.Set(name, value)
	return node
}

func (node HTMLNode) SetType(value string) HTMLNode {
	return node.Set("type", value)
}

func (node HTMLNode) SetValue(value any) HTMLNode {
	return node.Set("value", value)
}

func (node HTMLNode) String() string {
	return node.value.String()
}

func (node HTMLNode) Call(name string, args ...any) HTMLNode {
	return HTMLNode{value: node.value.Call(name, args...)}
}

func (node HTMLNode) Tailwind(classes ...tlw.TailwindClass) HTMLNode {
	for _, class := range classes {
		node.value.Get("classList").Call("add", string(class))
	}
	return node
}

func ToGlobalValue(value js.Value) HTMLNode {
	return HTMLNode{value: value}
}

func (node HTMLNode) AddRow(data []string) {
	row := TR()
	for _, cellData := range data {
		cell := TD().SetInnerHTML(cellData)
		row.Child(cell)
	}
	node.Child(row)
}

func (node HTMLNode) ReplaceWith(newNode HTMLNode) {
	node.value.Call("replaceWith", newNode.Value())
}
