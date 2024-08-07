package menu

import (
	"wasm/client/css"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
	"wasm/pkg/router"
)

type menuLink struct {
	Path string
	Name string
}

func Render(r *router.Router) dom.HTMLNode {
	sidebar := dom.Div().SetID("sidebar").Tailwind(
		tlw.Flex, tlw.FlexCol, tlw.HScreen, tlw.W64, tlw.Py4, tlw.Px3, tlw.SpaceY6).AddClass(css.BgTeriary700)

	linkStyles := []tlw.TailwindClass{tlw.Block, tlw.Py2, tlw.Px4, tlw.RoundedMd, tlw.TextGray300, tlw.HoverBgGray700, tlw.HoverTextWhite, tlw.CursorPointer}

	links := []menuLink{
		{Path: "/", Name: "Home"},
		{Path: "/todolist", Name: "Todo List"},
		{Path: "/counter", Name: "Counter"},
		{Path: "/counter-signal", Name: "Counter Signal"},
		{Path: "/user-management", Name: "User Management"},
		{Path: "/chat", Name: "Chat"},
		{Path: "/osb", Name: "OSB"},
		{Path: "/performance", Name: "Performance"},
	}

	for _, link := range links {
		l := dom.Anchor(link.Name).Tailwind(linkStyles...).OnClick(func(path string) dom.Func {
			return func(_ dom.Event) {
				r.NavigateTo(path)
			}
		}(link.Path))
		sidebar.Child(l)
	}

	return sidebar
}
