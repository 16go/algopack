package list

import (
	"github.com/16go/algopack"
	"github.com/16go/algopack/utils"
	"sync"
)

type node[T any] struct {
	data T
	next *node[T]
	prev *node[T]
}

type doubleLinkedList[T any] struct {
	mu    sync.Mutex
	safe  bool
	first *node[T]
	last  *node[T]
	size  int
}

func (l doubleLinkedList[T]) First() T {
	if l.size == 0 {
		return utils.Zero[T]()
	}
	return l.first.data
}

func (l doubleLinkedList[T]) Last() T {
	if l.size == 0 {
		return utils.Zero[T]()
	}
	return l.last.data
}

func (l doubleLinkedList[T]) Add(new *node[T]) {
	if l.safe {
		l.mu.Lock()
		defer l.mu.Unlock()
	}
	if l.first == nil {
		l.first = new
	} else {
		l.last.prev.next = new
		new.prev = l.last
		l.last = new
	}
	l.size++
}

func (l doubleLinkedList[T]) InsertBefore(target *node[T], new *node[T]) {
	if target == nil {
		utils.Panic("InsertBefore, target node is nil")
	}
	if l.safe {
		l.mu.Lock()
		defer l.mu.Unlock()
	}
	new.next = target
	if target.prev != nil {
		new.prev = target.prev
		target.prev = new
		target.prev.next = new
	} else {
		l.first = new
	}
}

func (l doubleLinkedList[T]) InsertAfter(target *node[T], new *node[T]) {
	if target == nil {
		utils.Panic("InsertAfter, target node is nil")
	}
	if l.safe {
		l.mu.Lock()
		defer l.mu.Unlock()
	}
}

func (l doubleLinkedList[T]) Iterate(fn algopack.IteratorFn, dir algopack.IteratorDirection) {
	var p *node[T]
	switch dir {
	case algopack.IteratorForward:
		p = l.first
	case algopack.IteratorBackward:
		p = l.last
	}
	for p != nil {
		stop := fn(p.data)
		if stop {
			return
		}
		// Move pointer to the next node.
		switch dir {
		case algopack.IteratorForward:
			p = p.next
		case algopack.IteratorBackward:
			p = p.prev
		}
	}
}
