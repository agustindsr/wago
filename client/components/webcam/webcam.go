package webcam

import (
	"syscall/js"
	"wasm/pkg/dom"
	"wasm/pkg/dom/tailwind"
)

type CameraComponent struct {
	videoElement dom.HTMLNode
	startButton  dom.HTMLNode
	stopButton   dom.HTMLNode
	stream       js.Value
}

func NewCameraComponent() *CameraComponent {
	cc := &CameraComponent{}

	cc.videoElement = dom.Document().Call("createElement", "video")
	cc.videoElement.Set("autoplay", true)
	cc.videoElement.Set("playsinline", true)
	cc.videoElement.Tailwind(tlw.WFull, tlw.HAuto, tlw.Border, tlw.BorderGray300, tlw.RoundedMd, tlw.H4)

	cc.startButton = dom.Button("Start Camera").
		Tailwind(tlw.P2, tlw.BgGreen500, tlw.TextWhite, tlw.RoundedMd, tlw.HoverBgGreen700).
		OnClick(cc.StartCamera)

	cc.stopButton = dom.Button("Stop Camera").
		Tailwind(tlw.P2, tlw.BgRed500, tlw.TextWhite, tlw.RoundedMd, tlw.HoverBgRed700).
		OnClick(cc.StopCamera)

	return cc
}

func (cc *CameraComponent) Render() dom.HTMLNode {
	container := dom.Div().Tailwind(tlw.MxAuto, tlw.P4, tlw.ShadowMd, tlw.RoundedLg).
		Child(
			dom.H2("Camera Test").Tailwind(tlw.Text2Xl, tlw.FontBold, tlw.Mb4),
			dom.Div().Tailwind(tlw.Flex, tlw.SpaceX2).
				Child(cc.startButton, cc.stopButton),
			cc.videoElement,
		)

	return container
}

func (cc *CameraComponent) StartCamera(_ dom.Event) {
	navigator := js.Global().Get("navigator")
	mediaDevices := navigator.Get("mediaDevices")
	if mediaDevices.Truthy() {
		constraints := map[string]interface{}{
			"video": true,
			"audio": false,
		}
		promise := mediaDevices.Call("getUserMedia", constraints)
		promise.Call("then", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			cc.stream = args[0]
			cc.videoElement.Set("srcObject", cc.stream)
			return nil
		}))
		promise.Call("catch", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			js.Global().Call("alert", "Error accessing camera: "+args[0].String())
			return nil
		}))
	} else {
		js.Global().Call("alert", "Media devices not supported")
	}
}

func (cc *CameraComponent) StopCamera(_ dom.Event) {
	if cc.stream.Truthy() {
		tracks := cc.stream.Call("getTracks")
		length := tracks.Length()
		for i := 0; i < length; i++ {
			track := tracks.Index(i)
			track.Call("stop")
		}
		cc.stream = js.Undefined()
		cc.videoElement.Set("srcObject", nil)
	}
}

func main() {
	document := dom.Document()
	body := document.Get("body")

	cc := NewCameraComponent()
	body.Call("appendChild", cc.Render())

	select {}
}
