## Benchmarking
The results for the interface and generic stack implementations.
```go
type ifaceStack struct {
    items []any
    top   int
    safe  bool
    mu    sync.Mutex
}

type Stack[T any] struct {
    items []T
    top   int
    safe  bool
    mu    sync.Mutex
}
```

| Benchmark | Iterations | Time per operation | Bytes allocated per operation | Allocations per operation |
| --- | --- | --- | --- | --- |
| BenchmarkIStack-12 | 65890224 | 58.44 ns/op | 101 B/op | 0 allocs/op |
| BenchmarkStack-12 | 425949592 | 8.733 ns/op | 41 B/op | 0 allocs/op |
