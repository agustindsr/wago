package dom

import "syscall/js"

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

type (
	NavElement struct {
		HTMLNode
	}

	FooterElement struct {
		HTMLNode
	}

	TableElement struct {
		HTMLNode
	}

	TRElement struct {
		HTMLNode
	}

	TDElement struct {
		HTMLNode
	}

	DivElement struct {
		HTMLNode
	}

	H1Element struct {
		HTMLNode
	}

	H2Element struct {
		HTMLNode
	}

	H3Element struct {
		HTMLNode
	}

	H4Element struct {
		HTMLNode
	}

	H5Element struct {
		HTMLNode
	}

	H6Element struct {
		HTMLNode
	}

	SpanElement struct {
		HTMLNode
	}

	PElement struct {
		HTMLNode
	}

	BRElement struct {
		HTMLNode
	}

	StrongElement struct {
		HTMLNode
	}

	BodyElement struct {
		HTMLNode
	}

	HeadElement struct {
		HTMLNode
	}
	AnchorElement struct {
		HTMLNode
	}

	InputElement struct {
		HTMLNode
	}

	ButtonElement struct {
		HTMLNode
	}

	LIElement struct {
		HTMLNode
	}

	ULElement struct {
		HTMLNode
	}
)

func Anchor() AnchorElement {
	return AnchorElement{Element("a")}
}

// Div crea un nuevo elemento div
func Div() DivElement {
	return DivElement{Element(divElement)}
}

func Span() SpanElement {
	return SpanElement{Element(spanElement)}
}

func P() PElement {
	return PElement{Element(pElement)}
}

func BR() BRElement {
	return BRElement{Element(brElement)}
}

func Strong() StrongElement {
	return StrongElement{Element(strongElement)}
}

func Table() TableElement {
	return TableElement{Element(tableElement)}
}

func TR() TRElement {
	return TRElement{Element(trElement)}
}

func TD() TDElement {
	return TDElement{Element(tdElement)}
}

func Footer() FooterElement {
	return FooterElement{Element(footerElement)}
}

func Nav() NavElement {
	return NavElement{Element(navElement)}
}

func H1() H1Element {
	return H1Element{Element(h1Element)}
}

func H2() H2Element {
	return H2Element{Element(h2Element)}
}

func H3() H3Element {
	return H3Element{Element(h3Element)}
}

func H4() H4Element {
	return H4Element{Element(h4Element)}
}

func H5() H5Element {
	return H5Element{Element(h5Element)}
}

func H6() H6Element {
	return H6Element{Element(h6Element)}
}

func Head() HeadElement {
	return HeadElement{Element("head")}
}

func Input() InputElement {
	return InputElement{Element(inputElement)}
}

func Button() ButtonElement {
	return ButtonElement{Element(buttonElement)}
}

func UL() ULElement {
	return ULElement{Element(ulElement)}
}

func LI() LIElement {
	return LIElement{Element(liElement)}
}

// SetInnerHTML establece el contenido HTML interno del elemento
func (gv HTMLNode) SetInnerHTML(html any) HTMLNode {
	gv.value.Set("innerHTML", html)
	return gv
}

// SetInnerText establece el texto interno del elemento
func (gv HTMLNode) SetInnerText(text string) HTMLNode {
	gv.value.Set("innerText", text)
	return gv
}

// Child agrega un elemento hijo
func (gv HTMLNode) Child(child HTMLNode) HTMLNode {
	gv.value.Call("appendChild", child.value)
	return gv
}

// ClearChildren elimina todos los hijos del elemento actual
func (gv HTMLNode) ClearChildren() HTMLNode {
	gv.SetInnerHTML("")
	return gv
}

func (gv HTMLNode) Get(name string) HTMLNode {
	return HTMLNode{value: gv.value.Get(name)}
}

func (gv HTMLNode) Set(name string, value any) HTMLNode {
	gv.value.Set(name, value)
	return gv
}

func (gv HTMLNode) String() string {
	return gv.value.String()
}

// Call invoca un método en el elemento
func (gv HTMLNode) Call(name string, args ...any) HTMLNode {
	return HTMLNode{value: gv.value.Call(name, args...)}
}

func ToGlobalValue(value js.Value) HTMLNode {
	return HTMLNode{value: value}
}

func (t TableElement) AddRow(data []string) {
	row := TR()
	for _, cellData := range data {
		cell := TD().SetInnerHTML(cellData)
		row.Child(cell)
	}
	t.Child(row.HTMLNode)
}
