package body

import (
	"wasm/components/todolist"
	"wasm/wa/dom"
)

func Render() dom.HTMLNode {
	container := dom.Div()
	container.AddClass("container")

	photosDiv := dom.Div()
	container.Child(photosDiv.HTMLNode)

	container.
		Child(dom.H1().SetInnerHTML("Welcome to My Website")).
		Child(dom.P().SetInnerHTML("This is a basic example of a website using WebAssembly with Go.")).
		Child(dom.P().SetInnerHTML("It includes a navigation bar, content area, and a footer."))

	/*
			photosTableCh := jsonapi.GetPhotosTable()

			go func() {
			photosTable := <-photosTableCh
			photosDiv.Child(photosTable.HTMLNode)
		}()

		dom.AppendToBody(container.HTMLNode)*/

	todoList := todolist.Render()

	container.Child(todoList.HTMLNode)

	return container.HTMLNode
}
