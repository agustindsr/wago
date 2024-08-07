package context

import (
	"context"
	"sync"
	"time"
)

type WagoCtx struct {
	mu    sync.RWMutex
	store map[interface{}]interface{}

	parent context.Context
	done   chan struct{}
	err    error
}

func Context(parent context.Context) *WagoCtx {
	return &WagoCtx{
		store:  make(map[interface{}]interface{}),
		parent: parent,
		done:   make(chan struct{}),
	}
}

func (ctx *WagoCtx) Get(key string) (interface{}, bool) {
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()
	value, exists := ctx.store[key]
	return value, exists
}

func (ctx *WagoCtx) GetString(key string) (string, bool) {
	value, exists := ctx.Get(key)
	if !exists {
		return "", false
	}
	str, ok := value.(string)
	return str, ok
}

func (ctx *WagoCtx) GetInt(key string) (int, bool) {
	value, exists := ctx.Get(key)
	if !exists {
		return 0, false
	}
	i, ok := value.(int)
	return i, ok
}

// GetBool recupera un valor booleano del contexto personalizado.
func (ctx *WagoCtx) GetBool(key string) (bool, bool) {
	value, exists := ctx.Get(key)
	if !exists {
		return false, false
	}
	b, ok := value.(bool)
	return b, ok
}

func (ctx *WagoCtx) Deadline() (time.Time, bool) {
	if ctx.parent != nil {
		return ctx.parent.Deadline()
	}
	return time.Time{}, false
}

func (ctx *WagoCtx) Done() <-chan struct{} {
	return ctx.done
}

func (ctx *WagoCtx) Err() error {
	return ctx.err
}

func (ctx *WagoCtx) Value(key interface{}) interface{} {
	ctx.mu.RLock()
	defer ctx.mu.RUnlock()
	if value, exists := ctx.store[key]; exists {
		return value
	}
	if ctx.parent != nil {
		return ctx.parent.Value(key)
	}
	return nil
}

func (ctx *WagoCtx) Set(key, value interface{}) {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	ctx.store[key] = value
}

func (ctx *WagoCtx) Cancel() {
	ctx.mu.Lock()
	defer ctx.mu.Unlock()
	if ctx.err == nil {
		ctx.err = context.Canceled
		close(ctx.done)
	}
}
