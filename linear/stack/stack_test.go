package stack_test

import (
	algo "github.com/16go/algopack"
	"github.com/16go/algopack/linear/stack"
	"sync"
	"testing"
)

var _ algo.ConcurrentSafeInterface = stack.NewStack[int]()
var _ algo.CollectionInterface = stack.NewStack[int]()
var _ algo.StackInterface[int] = stack.NewStack[int]()

func NewIfaceStack() *ifaceStack {
	return new(ifaceStack)
}

type ifaceStack struct {
	items []any
	top   int
	safe  bool
	mu    sync.Mutex
}

func (s *ifaceStack) Push(el any) {
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

type el struct {
	index int
}

func TestStack_Pop(t *testing.T) {
	s := stack.NewStack[el]()
	s.Push(el{1})
	s.Push(el{2})
	s.Push(el{3})
	var expected = []int{3, 2, 1}
	for i := 0; i < len(expected); i++ {
		if v := s.Pop(); v.index != expected[i] {
			t.Errorf("Expected %v, but got %v", expected[i], v)
		}
	}
}

func Benchmark_InterfaceStack_Push(b *testing.B) {
	is := NewIfaceStack()
	for i := 0; i < b.N; i++ {
		is.Push(i)
	}
}

func Benchmark_GenericStack_Push(b *testing.B) {
	gs := stack.NewStack[el]()
	for i := 0; i < b.N; i++ {
		gs.Push(el{i})
	}
}
