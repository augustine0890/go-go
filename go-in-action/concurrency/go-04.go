// This sample program demonstrates how to use the atomic package functions
// Store and Load to provide safe access to numeric types.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown is a flag to alert running goroutines to shutdown
	shutdown int64

	// wg is used to wait for the program to finish
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	// Safely flag it is time to shutdown
	fmt.Println("Shutting down...")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

// doWork simulates a goroutine performing work and checking the Shutdown flag to terminate early.
func doWork(name string) {
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s down\n", name)
			break
		}
	}
}
