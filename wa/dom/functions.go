package dom

import "syscall/js"

type (
	Func func(this js.Value, p []js.Value) any
)

func RegisterFunc(name string, fn Func) {
	js.Global().Set(name, js.FuncOf(fn))
}
