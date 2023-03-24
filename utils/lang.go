package utils

import "fmt"

func Zero[T any]() T {
	var zero T
	return zero
}

func Panic(msg string) {
	panic(fmt.Sprintf("algopack: %s", msg))
}
