package router

import (
	"fmt"
	"syscall/js"
	"wasm/pkg/dom"
)

type Route struct {
	Path      string
	Component func() dom.HTMLNode
}

type Router struct {
	routes map[string]Route
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]Route),
	}
}

func (r *Router) AddRoute(path string, component func() dom.HTMLNode) {
	r.routes[path] = Route{
		Path:      path,
		Component: component,
	}
}

func (r *Router) NavigateTo(path string) {
	route, exists := r.routes[path]
	if !exists {
		// Handle 404 - Not Found
		dom.ConsoleLog(path)
		dom.ElementByID("content").SetInnerHTML("Page not found")
		return
	}

	dom.ConsoleLog(fmt.Sprintf("Navigating to %s", path))
	dom.PushHistoryState(nil, "", "#"+path)

	dom.ElementByID("content").SetInnerHTML("")
	dom.ElementByID("content").Child(route.Component())
}

func (r *Router) HandleNavigation() {
	// Get current path from window location (simulate hash routing)
	path := js.Global().Get("location").Get("hash").String()
	if len(path) > 1 && path[0] == '#' {
		path = path[1:]
	}
	r.NavigateTo(path)
}

func (r *Router) Initialize() {
	r.HandleNavigation()
	dom.ConsoleLog("Router initialized")

	// Listen for hash change events to handle navigation
	js.Global().Get("window").Call("addEventListener", "hashchange", js.FuncOf(func(this js.Value, args []js.Value) any {
		r.HandleNavigation()
		return nil
	}))
}
