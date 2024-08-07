package auth

import (
	"wasm/pkg/context"
	"wasm/pkg/dom"
	tlw "wasm/pkg/dom/tailwind"
)

type Login struct {
	emailInput    dom.HTMLNode
	passwordInput dom.HTMLNode
	ctx           *context.WagoCtx
}

func New(ctx *context.WagoCtx) *Login {
	return &Login{
		ctx: ctx,
	}
}

func (l *Login) Render() dom.HTMLNode {
	l.emailInput = dom.Input().SetType("email").SetPlaceholder("Email")
	l.passwordInput = dom.Input().SetType("password").SetPlaceholder("Password")

	loginButton := dom.Button("Login").OnClick(l.login)

	form := dom.Div().Tailwind(tlw.Flex, tlw.FlexCol, tlw.SpaceY4).
		Child(
			l.emailInput,
			l.passwordInput,
			loginButton,
		)

	return dom.Div().Tailwind(tlw.Flex, tlw.ItemsCenter, tlw.JustifyCenter, tlw.HScreen).
		Child(form)
}

func (l *Login) login(_ dom.Event) {
	email := l.emailInput.GetValue().String()
	_ = l.passwordInput.GetValue().String()

	// Simulating a token after login
	token := "eyJraWQiOiJJeHAzSm1ETTdtMFNvYXpMWlk1UEluN1o5VVwvN3I2QWJUV1ArRkNyMkx5QT0iLCJhbGciOiJSUzI1NiJ9..."

	l.ctx.Set("token", token)
	l.ctx.Set("isAuth", true)
	l.ctx.Set("userID", "640b487883328a130f3ee799")
	l.ctx.Set("email", email)
	l.ctx.Set("name", "User")

}
