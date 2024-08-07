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

func (r *Router) HandleNavigation() {
	path := js.Global().Get("location").Get("hash").String()
	if len(path) > 1 && path[0] == '#' {
		path = path[1:]
	}
	r.NavigateTo(path)
}

func (r *Router) Initialize() {
	r.HandleNavigation()

	// Listen for hash change events to handle navigation
	js.Global().Get("window").Call("addEventListener", "hashchange", js.FuncOf(func(this js.Value, args []js.Value) any {
		dom.ConsoleLog("Hash change detected")
		r.HandleNavigation()
		return nil
	}))

	// Listen for popstate events to handle navigation
	js.Global().Get("window").Call("addEventListener", "popstate", js.FuncOf(func(this js.Value, args []js.Value) any {
		dom.ConsoleLog("Popstate detected")
		r.HandleNavigation()
		return nil
	}))

	// Handle the initial URL when the app is loaded
	r.HandleInitialURL()
}

func (r *Router) HandleInitialURL() {
	path := js.Global().Get("location").Get("pathname").String()
	if path == "" {
		path = "/"
	}
	r.NavigateTo(path)
}

func (r *Router) NavigateTo(path string) {
	route, exists := r.routes[path]
	if !exists {
		// Handle 404 - Not Found
		dom.ConsoleLog(fmt.Sprintf("Route not found: %s", path))
		dom.ConsoleLog(fmt.Sprintf("routers: %v", r.routes))
		dom.ElementByID("content").SetInnerHTML("Page not found")
		return
	}

	dom.ConsoleLog(fmt.Sprintf("Navigating to %s", path))
	dom.PushHistoryState(nil, "", path)

	dom.ElementByID("content").SetInnerHTML("")
	dom.ElementByID("content").Child(route.Component())
}
