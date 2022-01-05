// This sample program demonstrates how to use a mutex to define critical sections of code that need synchronous access.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines
	counter int

	wg sync.WaitGroup

	// mutex is used to define a critical section of code
	mutex sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Printf("Final Count: %d\n", counter)
}

// incCounter increments the package level Counter variable
// using the Mutex to synchronize and provide safe access.
func incCounter(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		// Only allow one goroutine through this critical section at a time
		mutex.Lock()
		{
			// Capture the value of counter
			value := counter

			// Yield the thread and be placed back in queue
			runtime.Gosched()

			// Increment our local value of counter
			value++
			// Store the value back to counter
			counter = value
		}
		// Release the lock and allow any waiting goroutines through
		mutex.Unlock()
	}
}
