package utils_test

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Set the maximum number of threads to the number of CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create a WaitGroup to wait for all goroutines to finish
	wg := sync.WaitGroup{}

	// Spawn a goroutine for each CPU core
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			// Perform some work in parallel
			fmt.Println("Hello from goroutine!")
			for i := 0; i < 10_000_000_000; i++ {
				_ = i * i
			}
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
