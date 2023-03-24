package algopack

type (
	IteratorDirection int
	// IteratorFn is being called each time an element of the collection is being iterated over.
	// Use a closure to pass extra arguments to the iterator function.
	// For instance:
	// `func (arg1 int, arg2 string) IteratorFn {
	// 		return func(data any) bool {
	//			// Do something with item data
	//			// ...
	//			return false
	//		}
	// }`
	IteratorFn func(data any) (stop bool)
)

const (
	IteratorForward IteratorDirection = iota
	IteratorBackward
)

type ConcurrentSafeInterface interface {
	EnableConcurrency()
	IsEnabled() bool
}

type CollectionInterface interface {
	IsEmpty() bool
	Size() int
}

type IteratorInterface[T any] interface {
	ConcurrentSafeInterface
	Next() T
	Rewind()
}

type StackInterface[T any] interface {
	CollectionInterface
	ConcurrentSafeInterface
	Push(T)
	Pop() T
	Peek() T
}
