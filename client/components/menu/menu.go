package menu

import (
	"wasm/client/components/chat"
	"wasm/client/components/counter"
	"wasm/client/components/draftea/osb"
	"wasm/client/components/home"
	"wasm/client/components/todolist"
	"wasm/client/components/usermanagement"
	"wasm/client/css"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

func Render() dom.HTMLNode {
	sidebar := dom.Div().SetID("sidebar").Tailwind(
		tlw.Flex, tlw.FlexCol, tlw.HScreen, tlw.W64, tlw.Py4, tlw.Px3, tlw.SpaceY6).AddClass(css.BgTeriary500)

	linkStyles := []tlw.TailwindClass{tlw.Block, tlw.Py2, tlw.Px4, tlw.RoundedMd, tlw.TextGray300, tlw.HoverBgGray700, tlw.HoverTextWhite, tlw.CursorPointer}

	homeLink := dom.Anchor("Home").Tailwind(linkStyles...).
		OnClick(navigateTo(home.Render()))
	todoLink := dom.Anchor("Todo List").Tailwind(linkStyles...).
		OnClick(navigateTo(todolist.Render()))
	counterLink := dom.Anchor("Counter").Tailwind(linkStyles...).
		OnClick(navigateTo(counter.NewCounter().Render()))

	counterSignalTC39Link := dom.Anchor("Counter Signal").Tailwind(linkStyles...).
		OnClick(navigateTo(counter.NewCounterSignal().Render()))

	userManagementLink := dom.Anchor("User Management").Tailwind(linkStyles...).
		OnClick(navigateTo(usermanagement.Render()))
	chatLink := dom.Anchor("Chat").Tailwind(linkStyles...).
		OnClick(navigateTo(chat.New().Render()))

	betsLink := dom.Anchor("OSB").Tailwind(linkStyles...).
		OnClick(navigateTo(osb.NewOSB().Render()))

	minimizeButton := dom.Button("Minimize").Tailwind(tlw.P2, tlw.BgRed500, tlw.TextWhite, tlw.Rounded, tlw.CursorPointer).
		OnClick(toggleSidebar)

	sidebar.Child(minimizeButton, homeLink, counterLink, counterSignalTC39Link, todoLink, userManagementLink, chatLink, betsLink)

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

func navigateTo(component dom.HTMLNode) dom.Func {
	return func(_ dom.Event) {
		dom.ElementByID("content").SetInnerHTML("")
		dom.ElementByID("content").Child(component)
	}
}
