package modal

import (
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
)

type Size int

const (
	Small Size = iota
	Medium
	Large
)

var modalSizeClasses = map[Size][]tlw.TailwindClass{
	Small:  {tlw.MaxWSm, tlw.WFull},
	Medium: {tlw.MaxW3Xl, tlw.WFull},
	Large:  {tlw.MaxW6Xl, tlw.WFull},
}

type ModalOption func(*Modal)

type Modal struct {
	Node     dom.HTMLNode
	CloseBtn dom.HTMLNode
	Content  dom.HTMLNode
	OnClose  dom.Func
	Title    string
	Size     []tlw.TailwindClass
}

func WithContent(content dom.HTMLNode) ModalOption {
	return func(m *Modal) {
		m.Content = content
	}
}

func WithOnClose(onClose dom.Func) ModalOption {
	return func(m *Modal) {
		m.OnClose = onClose
	}
}

func WithTitle(title string) ModalOption {
	return func(m *Modal) {
		m.Title = title
	}
}

func WithSize(size Size) ModalOption {
	return func(m *Modal) {
		m.Size = modalSizeClasses[size]
	}
}

func NewModal(opts ...ModalOption) *Modal {
	modal := &Modal{}

	modal.CloseBtn = dom.Button("x").Tailwind(
		tlw.Absolute, tlw.Top0, tlw.Right0, tlw.M4, tlw.P2, tlw.BgTransparent, tlw.TextGray700, tlw.HoverTextBlack,
	).OnClick(func(_ dom.Event) {
		modal.Close()
	})

	for _, opt := range opts {
		opt(modal)
	}

	modalContent := dom.Div().Tailwind(
		tlw.BgWhite, tlw.P6, tlw.RoundedLg, tlw.ShadowLg, tlw.WFull, tlw.MaxWSm, tlw.MxAuto, tlw.Relative,
	).Child(
		modal.CloseBtn,
	)

	if modal.Title != "" {
		titleElem := dom.H2(modal.Title).Tailwind(tlw.TextXl, tlw.FontBold, tlw.Mb4)
		modalContent.Child(titleElem)
	}

	modalContent.Child(modal.Content)

	if len(modal.Size) > 0 {
		modalContent.Tailwind(modal.Size...)
	}

	modal.Node = dom.Div().Tailwind(
		tlw.Fixed, tlw.Inset0, tlw.Flex, tlw.ItemsCenter, tlw.JustifyCenter, tlw.BgGray800, tlw.BgOpacity50,
	).Child(
		modalContent,
	)

	return modal
}

func (m *Modal) Close() {
	m.Node.Call("remove")
}

func (m *Modal) Open() {
	dom.Body().Child(m.Node)
}
