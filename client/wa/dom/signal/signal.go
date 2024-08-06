package signal

import "sync"

// SignalInterface defines the methods that both Signal and ComputedSignal must implement
type SignalInterface interface {
	Get() any
	Subscribe(listener func(any))
}

// Signal represents a simple signal
type Signal[T any] struct {
	mu        sync.RWMutex
	value     T
	listeners []func(T)
}

func NewSignal[T any](initial T) *Signal[T] {
	return &Signal[T]{
		value:     initial,
		listeners: []func(T){},
	}
}

func (s *Signal[T]) Get() T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.value
}

func (s *Signal[T]) Set(newValue T) {
	s.mu.Lock()
	s.value = newValue
	listeners := s.listeners
	s.mu.Unlock()
	for _, listener := range listeners {
		listener(newValue)
	}
}

func (s *Signal[T]) Subscribe(listener func(T)) {
	s.mu.Lock()
	s.listeners = append(s.listeners, listener)
	s.mu.Unlock()
	listener(s.value) // Notify immediately with the current value
}

// ToSignalInterface converts a Signal to SignalInterface
func (s *Signal[T]) ToSignalInterface() SignalInterface {
	return &signalWrapper[T]{s}
}

// signalWrapper wraps a Signal to implement SignalInterface
type signalWrapper[T any] struct {
	signal *Signal[T]
}

func (sw *signalWrapper[T]) Get() any {
	return sw.signal.Get()
}

func (sw *signalWrapper[T]) Subscribe(listener func(any)) {
	sw.signal.Subscribe(func(value T) {
		listener(value)
	})
}
