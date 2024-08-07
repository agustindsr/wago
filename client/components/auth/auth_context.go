package auth

import (
	"syscall/js"
	"wasm/pkg/context"
	"wasm/pkg/dom"
)

type Context struct {
	Context *context.WagoCtx
}

func NewAuthContext(ctx *context.WagoCtx) *Context {
	authCtx := &Context{Context: ctx}
	authCtx.CheckAuth()
	return authCtx
}

func (authCtx *Context) CheckAuth() {
	token := js.Global().Get("localStorage").Call("getItem", "token").String()
	if token != "" {
		authCtx.Context.Set("token", token)
		authCtx.Context.Set("isAuth", true)
		authCtx.Context.Set("userID", "640b487883328a130f3ee799")
		authCtx.Context.Set("email", "agustin.desautu+350@draftea.com")
		authCtx.Context.Set("name", "agustin")
	} else {
		authCtx.Context.Set("isAuth", false)
	}

	isAuth, _ := authCtx.Context.GetBool("isAuth")
	if !isAuth {
		authCtx.RedirectToLogin()
	}
}

func (authCtx *Context) Login(token string) {
	js.Global().Get("localStorage").Call("setItem", "token", token)
	authCtx.Context.Set("token", token)
	authCtx.Context.Set("isAuth", true)
	authCtx.Context.Set("userID", "640b487883328a130f3ee799")
	authCtx.Context.Set("email", "agustin.desautu+350@draftea.com")
	authCtx.Context.Set("name", "agustin")
}

func (authCtx *Context) Logout() {
	js.Global().Get("localStorage").Call("removeItem", "token")
	authCtx.Context.Set("token", "")
	authCtx.Context.Set("isAuth", false)
	authCtx.Context.Set("userID", "")
	authCtx.Context.Set("email", "")
	authCtx.Context.Set("name", "")
	authCtx.RedirectToLogin()
}

func (authCtx *Context) RedirectToLogin() {
	dom.ElementByID("app").SetInnerHTML("")
	dom.ElementByID("app").Child(New(authCtx.Context).Render())
}
