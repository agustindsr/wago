package signal

import "sync"

// ComputedSignal represents a computed signal
type ComputedSignal[T any] struct {
	mu        sync.RWMutex
	value     T
	compute   func() T
	listener  func()
	listeners []func(T)
	depends   []SignalInterface
}

func NewComputedSignal[T any](compute func() T, signals ...SignalInterface) *ComputedSignal[T] {
	cs := &ComputedSignal[T]{
		compute: compute,
	}
	cs.listener = func() {
		cs.recompute()
	}
	for _, s := range signals {
		s.Subscribe(func(any) {
			cs.listener()
		})
		cs.depends = append(cs.depends, s)
	}
	cs.recompute()
	return cs
}

func (cs *ComputedSignal[T]) Get() T {
	cs.mu.RLock()
	defer cs.mu.RUnlock()
	return cs.value
}

func (cs *ComputedSignal[T]) recompute() {
	cs.mu.Lock()
	cs.value = cs.compute()
	listeners := cs.listeners
	cs.mu.Unlock()
	for _, listener := range listeners {
		listener(cs.value)
	}
}

func (cs *ComputedSignal[T]) Subscribe(listener func(T)) {
	cs.mu.Lock()
	cs.listeners = append(cs.listeners, listener)
	cs.mu.Unlock()
	listener(cs.value) // Notify immediately with the current value
}

// ToSignalInterface converts a ComputedSignal to SignalInterface
func (cs *ComputedSignal[T]) ToSignalInterface() SignalInterface {
	return &computedWrapper[T]{cs}
}

// computedWrapper wraps a ComputedSignal to implement SignalInterface
type computedWrapper[T any] struct {
	computed *ComputedSignal[T]
}

func (cw *computedWrapper[T]) Get() any {
	return cw.computed.Get()
}

func (cw *computedWrapper[T]) Subscribe(listener func(any)) {
	cw.computed.Subscribe(func(value T) {
		listener(value)
	})
}
