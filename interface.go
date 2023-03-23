package main

type StackInterface[T any] interface {
	Push(T)
	Pop() T
	IsEmpty() bool
	Peek() T
}
