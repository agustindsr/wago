package home

import (
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
)

func Render() dom.HTMLNode {
	container := dom.Div().Tailwind(tlw.P4)

	title := dom.H1("Welcome to the Home Page").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4)
	description := dom.P("WebAssembly (Wasm) is a binary instruction format for a stack-based virtual machine. " +
		"It is designed as a portable compilation target for programming languages, enabling deployment on the web for client and server applications.").Tailwind(tlw.Mb4)

	subtitle := dom.H2("Using WebAssembly with Go").Tailwind(tlw.TextXl, tlw.FontBold, tlw.Mb4)
	content := dom.P("Go can be compiled to WebAssembly, allowing Go code to run in web browsers and other environments that support WebAssembly. " +
		"This makes it possible to write web applications entirely in Go, sharing code between the frontend and backend.").Tailwind(tlw.Mb4)

	stepsTitle := dom.H2("Steps to Use WebAssembly with Go").Tailwind(tlw.TextXl, tlw.FontBold, tlw.Mb4)
	steps := dom.P("1. Write your Go code.\n" +
		"2. Compile the Go code to WebAssembly using the Go compiler.\n" +
		"3. Load the WebAssembly module in your web application.\n" +
		"4. Interact with the WebAssembly module from JavaScript.").Tailwind(tlw.Mb4)

	exampleTitle := dom.H2("Example").Tailwind(tlw.TextXl, tlw.FontBold, tlw.Mb4)
	example := dom.P("Here is a simple example of a Go function compiled to WebAssembly and called from JavaScript:").Tailwind(tlw.Mb4)
	codeExample := dom.Pre(`
<code class="block p-2 bg-gray-200 rounded">
package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("add", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		a := p[0].Int()
		b := p[1].Int()
		return a + b
	}))
	select {}
}
</code>`).Tailwind(tlw.Block, tlw.P2, tlw.BgGray200, tlw.Rounded)

	container.Child(
		title,
		description,
		subtitle,
		content,
		stepsTitle,
		steps,
		exampleTitle,
		example,
		codeExample,
	)

	return container
}
