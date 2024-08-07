package menu

import (
	"wasm/client/css"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
	"wasm/pkg/router"
)

func Render(r *router.Router) dom.HTMLNode {
	sidebar := dom.Div().SetID("sidebar").Tailwind(
		tlw.Flex, tlw.FlexCol, tlw.HScreen, tlw.W64, tlw.Py4, tlw.Px3, tlw.SpaceY6).AddClass(css.BgTeriary500)

	linkStyles := []tlw.TailwindClass{tlw.Block, tlw.Py2, tlw.Px4, tlw.RoundedMd, tlw.TextGray300, tlw.HoverBgGray700, tlw.HoverTextWhite, tlw.CursorPointer}

	homeLink := dom.Anchor("Home").Tailwind(linkStyles...).OnClick(func(_ dom.Event) {
		r.NavigateTo("/")
	})

	todoLink := dom.Anchor("Todo List").Tailwind(linkStyles...).OnClick(func(_ dom.Event) {
		r.NavigateTo("/todolist")
	})
	counterLink := dom.Anchor("Counter").Tailwind(linkStyles...).OnClick(func(_ dom.Event) {
		r.NavigateTo("/counter")
	})
	counterSignalTC39Link := dom.Anchor("Counter Signal").Tailwind(linkStyles...).OnClick(func(_ dom.Event) {
		r.NavigateTo("/counter-signal")
	})
	userManagementLink := dom.Anchor("User Management").Tailwind(linkStyles...).OnClick(func(_ dom.Event) {
		r.NavigateTo("/usermanagement")
	})
	chatLink := dom.Anchor("Chat").Tailwind(linkStyles...).OnClick(func(_ dom.Event) {
		r.NavigateTo("/chat")
	})
	betsLink := dom.Anchor("OSB").Tailwind(linkStyles...).OnClick(func(_ dom.Event) {
		r.NavigateTo("/osb")
	})

	minimizeButton := dom.Button("Minimize").Tailwind(tlw.P2, tlw.BgRed500, tlw.TextWhite, tlw.Rounded, tlw.CursorPointer).OnClick(toggleSidebar)

	sidebar.Child(minimizeButton, homeLink, todoLink, counterLink, counterSignalTC39Link, userManagementLink, chatLink, betsLink)

	return sidebar
}

func toggleSidebar(_ dom.Event) {
	sidebar := dom.ElementByID("sidebar")
	if sidebar.HasClass(css.SidebarCollapsed) {
		sidebar.RemoveClass(css.SidebarCollapsed)
	} else {
		sidebar.AddClass(css.SidebarCollapsed)
	}
}
