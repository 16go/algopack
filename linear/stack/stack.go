package stack

import (
	"github.com/16go/algopack/utils"
	"sync"
)

func NewStack[T any]() *stack[T] {
	return new(stack[T])
}

// NewSafeStack creates a concurrent-safe stack.
func NewSafeStack[T any]() *stack[T] {
	s := new(stack[T])
	s.safe = true
	return s
}

type stack[T any] struct {
	items []T
	top   int
	safe  bool
	mu    sync.Mutex
}

func (s *stack[T]) IsEnabled() bool {
	return s.safe
}

func (s *stack[T]) EnableConcurrency() {
	s.safe = true
}

func (s *stack[T]) Push(el T) {
	if s.safe {
		s.mu.Lock()
		defer s.mu.Unlock()
	}
	if s.top == len(s.items) {
		s.items = append(s.items, el)
	} else {
		s.items[s.top] = el
	}
	s.top++
}

func (s *stack[T]) Pop() T {
	if s.safe {
		s.mu.Lock()
		defer s.mu.Unlock()
	}
	if s.IsEmpty() {
		return utils.Zero[T]()
	}
	el := s.items[s.top-1]
	s.top--
	return el
}

// Peek returns an element from the stack without removing it.
// If the stack is empty, Peek returns nil.
func (s *stack[T]) Peek() T {
	if s.safe {
		s.mu.Lock()
		defer s.mu.Unlock()
	}
	if s.IsEmpty() {
		return utils.Zero[T]()
	}
	return s.items[s.top]
}

// IsEmpty returns TRUE if the stack does not have elements.
func (s *stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack.
func (s *stack[T]) Size() int {
	return len(s.items)
}
