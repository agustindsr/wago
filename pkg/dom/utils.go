package dom

import "syscall/js"

// Alert muestra una alerta en el navegador con el mensaje proporcionado
func Alert(message string) {
	js.Global().Call("alert", message)
}

// Global obtiene un valor global de JavaScript
func Global(name string) HTMLNode {
	return HTMLNode{value: js.Global().Get(name)}
}

// NewGlobalValue crea una nueva instancia de HTMLNode
func NewGlobalValue(name string) HTMLNode {
	return HTMLNode{value: js.Global().Get(name)}
}

func ConsoleLog(message any) {
	js.Global().Get("console").Call("log", message)
}
