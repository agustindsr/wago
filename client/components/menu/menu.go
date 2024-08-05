package menu

import (
	"wasm/client/components/chat"
	"wasm/client/components/counter"
	"wasm/client/components/home"
	"wasm/client/components/todolist"
	"wasm/client/components/usermanagement"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

func Render() dom.HTMLNode {
	sidebar := dom.Div().Tailwind(
		tlw.Flex, tlw.FlexCol, tlw.HScreen, tlw.W64, tlw.BgGray800, tlw.Py4, tlw.Px3, tlw.SpaceY6)

	linkStyles := []tlw.TailwindClass{tlw.Block, tlw.Py2, tlw.Px4, tlw.RoundedMd, tlw.TextGray300, tlw.HoverBgGray700, tlw.HoverTextWhite, tlw.CursorPointer}

	homeLink := dom.Anchor("Home").Tailwind(linkStyles...).
		OnClick(navigateTo(home.Render()))
	todoLink := dom.Anchor("Todo List").Tailwind(linkStyles...).
		OnClick(navigateTo(todolist.Render()))
	counterLink := dom.Anchor("Counter").Tailwind(linkStyles...).
		OnClick(navigateTo(counter.NewCounter().Render()))
	userManagementLink := dom.Anchor("User Management").Tailwind(linkStyles...).
		OnClick(navigateTo(usermanagement.Render()))
	chatLink := dom.Anchor("Chat").Tailwind(linkStyles...).
		OnClick(navigateTo(chat.New().Render()))

	betsLink := dom.Anchor("Bets").Tailwind(linkStyles...).
		OnClick(navigateTo(home.Render()))

	sidebar.Child(homeLink, counterLink, todoLink, userManagementLink, chatLink, betsLink)

	return sidebar
}

func navigateTo(component dom.HTMLNode) dom.Func {
	return func(_ dom.Event) {
		dom.ElementByID("content").SetInnerHTML("")
		dom.ElementByID("content").Child(component)
	}
}
