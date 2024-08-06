package signal

import (
	"sync"
)

type Updatable interface {
	UpdateDOM()
}

type Signal[T any] struct {
	value     T
	listeners map[Updatable]func(T)
	mu        sync.Mutex
}

func NewSignal[T any](initial T) *Signal[T] {
	return &Signal[T]{
		value:     initial,
		listeners: make(map[Updatable]func(T)),
	}
}

func (s *Signal[T]) Get() T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.value
}

func (s *Signal[T]) Set(newValue T) {
	s.mu.Lock()
	s.value = newValue
	s.mu.Unlock()
	s.notify()
}

func (s *Signal[T]) Subscribe(listener func(T)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	listener(s.value) // Notify immediately with the current value
}

func (s *Signal[T]) notify() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for component, listener := range s.listeners {
		listener(s.value)
		component.UpdateDOM()
	}
}
