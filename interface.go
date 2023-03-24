package algopack

type ConcurrentSafeInterface interface {
	EnableConcurrency()
	IsEnabled() bool
}

type CollectionInterface interface {
	IsEmpty() bool
	Size() int
}

type StackInterface[T any] interface {
	CollectionInterface
	ConcurrentSafeInterface
	Push(T)
	Pop() T
	Peek() T
}
