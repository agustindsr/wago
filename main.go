//go:build js && wasm

package main

import (
	"wasm/client/components/home"
	"wasm/client/components/menu"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

func main() {
	dom.Body().Tailwind(tlw.Flex, tlw.HFull).
		Child(
			menu.Render(),
			dom.Div().
				SetID("content").
				Child(home.Render()).Tailwind(tlw.Flex1, tlw.P4, tlw.Mb8),
		)

	select {}
}
