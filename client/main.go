//go:build js && wasm

package main

import (
	"wasm/client/components/chat"
	"wasm/client/components/counter"
	"wasm/client/components/draftea/osb"
	"wasm/client/components/home"
	"wasm/client/components/menu"
	"wasm/client/components/todolist"
	"wasm/client/components/usermanagement"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
	"wasm/pkg/router"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			dom.ConsoleLog(r)
		}
	}()

	r := router.NewRouter()

	r.AddRoute("/", home.Render)
	r.AddRoute("/osb", osb.NewOSB().Render)
	r.AddRoute("/todolist", todolist.Render)
	r.AddRoute("/counter", counter.NewCounter().Render)
	r.AddRoute("/counter-signal", counter.NewCounterSignal().Render)
	r.AddRoute("/usermanagement", usermanagement.Render)
	r.AddRoute("/chat", chat.New().Render)

	dom.ElementByID("app").SetInnerHTML("")
	dom.ElementByID("app").Child(renderApp(r))

	// Initialize router to handle navigation
	r.Initialize()

	select {}
}

func renderApp(r *router.Router) dom.HTMLNode {
	container := dom.Div().Tailwind(tlw.Flex, tlw.HScreen).
		Child(
			menu.Render(r),
			dom.Div().SetID("content").Tailwind(tlw.Flex1, tlw.OverflowYAuto),
		)
	return container
}
