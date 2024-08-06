//go:build js && wasm

package main

import (
	"wasm/client/components/home"
	"wasm/client/components/menu"
	"wasm/client/css"
	"wasm/client/wa/dom"
	tlw "wasm/client/wa/dom/tailwind"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			dom.ConsoleLog(r)
		}
	}()
	dom.Body().Tailwind(tlw.Flex, tlw.HFull, tlw.TextWhite).AddClass(css.BgTeriary900).
		Child(
			menu.Render(),
			dom.Div().
				SetID("content").
				Child(home.Render()).Tailwind(tlw.Flex1, tlw.P4, tlw.Mb8),
		)

	select {}
}
