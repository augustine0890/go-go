// This sample program demonstrates how to create race conditions in our programs.
// We don't want to do this.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter is a variable incremented by all goroutines
	counter int64
	// wg is used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	// Create two goroutines
	go incCounter(1)
	go incCounter(2)

	// Wait for the goroutine to finish
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable
func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Capture the value of Counter
		// value := counter
		// Only one goroutine can perform and complete add operation at a time
		atomic.AddInt64(&counter, 1)
		// Yield the thread and be placed back in queue
		runtime.Gosched()

		// Increment out local value of Counter
		// value++

		// Store the value back into Counter`
		// counter = value
	}
}
