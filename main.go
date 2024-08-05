//go:build js && wasm

package main

import (
	"wasm/components/shared/body"
	"wasm/components/shared/footer"
	"wasm/components/shared/navbar"
)

func main() {
	navbar.Render()
	body.Render()

	footer.Render()

	select {}
}
